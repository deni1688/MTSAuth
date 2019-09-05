package api

import (
	"bytes"
	"encoding/json"
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

func TestCheckServiceController(t *testing.T) {
	expected, _ := json.Marshal(map[string]string{"status": "Service is running"})
	w := testRequest(t, "GET", "localhost:9000", "", c.CheckServiceController)

	assert.Equal(t, http.StatusAccepted, w.Code, "should equal 200")
	assert.Equal(t, string(expected), w.Body.String(), "should return Service is running")
}

func TestLoginControllerSuccess(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "testing123"}`
	expected, _ := json.Marshal(map[string]string{"token": "mockToken123"})

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusAccepted, w.Code, "should equal 200")
	assert.Equal(t, string(expected), w.Body.String(), "should return the token")
}

func TestLoginControllerValidationFails(t *testing.T) {
	data := `{"email": "", "password": ""}`
	expected, _ := json.Marshal(map[string]string{"error": "Invalid credentials format"})

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusBadRequest, w.Code, "should equal 400")
	assert.Equal(t, string(expected), w.Body.String(), "should return error for invalid email and password")
}

func TestLoginControllerAuthFails(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "wrongpass123"}`
	expected, _ := json.Marshal(map[string]string{"error": "Access Denied"})

	w := testRequest(t, "POST", "localhost:9000/login", data, c.LoginController)

	assert.Equal(t, http.StatusForbidden, w.Code, "should equal 403")
	assert.Equal(t, string(expected), w.Body.String(), "should return error for unauthenticated user")
}
