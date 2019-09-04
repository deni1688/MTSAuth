package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deni1688/motusauth/app"
)

var (
	domain = new(app.MockApp)
	c      = &Controller{domain}
)

func TestCheckServiceController(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9000", nil)

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	controller := http.HandlerFunc(c.CheckServiceController)

	controller.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	expected := map[string]string{"status": "Service is running"}
	expectedBytes, _ := json.Marshal(expected)

	if w.Body.String() != string(expectedBytes) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestLoginController(t *testing.T) {
	data := `{"email": "t@testing.com", "password": "testing123"}`

	req, err := http.NewRequest("POST", "localhost:9000/login", bytes.NewBufferString(data))

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	w := httptest.NewRecorder()

	controller := http.HandlerFunc(c.LoginController)

	controller.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	expected := map[string]string{"token": "mockToken123"}
	expectedBytes, _ := json.Marshal(expected)

	if w.Body.String() != string(expectedBytes) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expectedBytes))
	}
}
