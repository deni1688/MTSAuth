package api

import (
	"encoding/json"
	"net/http"

	"github.com/deni1688/motusauth/app"
	"github.com/deni1688/motusauth/models"
)

func CheckServiceController(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil || u.Password == "" || u.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials format")
		return
	}

	token, err := app.AuthenticateUser(&u)

	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]string{"token": token})
}

func SignUpController(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user format")
		return
	}

	user, err := app.CreateUser(&u)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusAccepted, user)
}
