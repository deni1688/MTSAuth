package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Credentials - username and password from the request body
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims - payload of the token
type Claims struct {
	FirstName string   `json:"frstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func handleServiceCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	var auth User

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil || auth.Password == "" || auth.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials format")
		return
	}

	user, err := auth.authenticateUser()

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Authentication failed")
		return
	}

	token, err := user.createTokenString()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]string{"token": token})
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register ")
}

func handleCreateManager(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register ")
}

func (user *User) createTokenString() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	standardClaims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	claims := &Claims{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Roles:          user.Roles,
		StandardClaims: standardClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
