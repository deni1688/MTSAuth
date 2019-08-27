package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims - payload returned to the client
type Claims struct {
	FirstName string   `json:"frstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

func handleServiceCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var reqUser User

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil || reqUser.Password == "" || reqUser.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials format")

		return
	}

	user, err := reqUser.authenticate()

	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())

		return
	}

	token, err := user.createTokenString()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]string{"token": token})
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	var reqUser User

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil ||
		reqUser.Password == "" ||
		reqUser.Email == "" ||
		reqUser.FirstName == "" ||
		reqUser.LastName == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid user format")

		return
	}

	companyID := hashAndSalt(reqUser.Email)
	password := hashAndSalt(reqUser.Password)

	reqUser.CompanyID = companyID
	reqUser.Password = password
	reqUser.CreatedAt = time.Now()
	reqUser.Roles = "ProductManager,OrderManager,UserManager"
	reqUser.IsRoot = true

	user := reqUser.createUser()

	respondWithJSON(w, http.StatusAccepted, user)
}
