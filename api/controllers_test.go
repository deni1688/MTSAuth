package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deni1688/motusauth/app"
	"github.com/stretchr/testify/assert"
)

var (
	domain = new(app.MockApp)
	c      = &Controller{domain}
)

func testRequest(t *testing.T, method string, url string, data string, c func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(data))

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	w := httptest.NewRecorder()

	controller := http.HandlerFunc(c)
	controller.ServeHTTP(w, req)

	return w
}

func Test_CheckServiceController_Success(t *testing.T) {
	expected := `{"status":"Service is running"}`
	w := testRequest(t, "GET", "localhost:9000", "", c.CheckServiceController)

	assert.Equal(t, http.StatusAccepted, w.Code, "should equal 200")
	assert.Equal(t, expected, w.Body.String(), "should return Service is running")
}

func Test_LoginController_Success(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "testing123"}`
	expected := `{"token":"mockToken123"}`

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusAccepted, w.Code, "should equal 200")
	assert.Equal(t, expected, w.Body.String(), "should return the token")
}

func Test_LoginController_ValidationFails(t *testing.T) {
	data := `{"email": "", "password": ""}`
	expected := `{"error":"Invalid credentials format"}`

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusBadRequest, w.Code, "should equal 400")
	assert.Equal(t, expected, w.Body.String(), "should return error for invalid email and password")
}

func Test_LoginController_AuthFails(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "wrongpass123"}`
	expected := `{"error":"Access Denied"}`

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusForbidden, w.Code, "should equal 403")
	assert.Equal(t, expected, w.Body.String(), "should return error for unauthenticated user")
}

func Test_RegisterController_Success(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "12345678", "firstname": "Mike", "lastname": "Jones"}`
	expected := `{"message":"User Mike Jones created"}`

	w := testRequest(t, "POST", "localhost:9000/register", data, c.RegisterController)

	assert.Equal(t, http.StatusAccepted, w.Code, "should equal 200")
	assert.Equal(t, expected, w.Body.String(), "should register a new user")
}

func Test_RegisterController_DecodeFails(t *testing.T) {
	data := ""
	expected := `{"error":"Invalid user format"}`

	w := testRequest(t, "POST", "localhost:9000/register", data, c.RegisterController)

	assert.Equal(t, http.StatusBadRequest, w.Code, "should equal 400")
	assert.Equal(t, expected, w.Body.String(), "should return error for bad json")
}

func Test_RegisterController_RegisterFails(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "1234567", "firstname": "Mike", "lastname": "Jones"}`
	expected := `{"error":"Password min 8 letters"}`

	w := testRequest(t, "POST", "localhost:9000/register", data, c.RegisterController)

	assert.Equal(t, http.StatusBadRequest, w.Code, "should equal 400")
	assert.Equal(t, expected, w.Body.String(), "should return error for invalid user")
}
