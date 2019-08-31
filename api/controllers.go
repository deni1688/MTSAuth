package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deni1688/motusauth/app"
	"github.com/deni1688/motusauth/models"
)

// CheckServiceController check tha the api is running
func CheckServiceController(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

// LoginController authentictes user by email and password
// and returns a signed jwt token
func LoginController(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials format")
		return
	}

	token, err := app.AuthenticateUser(u)

	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]string{"token": token})
}

// RegisterController handles the creation of a new root user
func RegisterController(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user format")
		return
	}

	if err := app.RegisterUser(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]string{"message": fmt.Sprintf("User %s %s created", u.FirstName, u.LastName)})
}