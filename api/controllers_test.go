package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckServiceController(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9000", nil)

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	controller := http.HandlerFunc(CheckServiceController)

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
