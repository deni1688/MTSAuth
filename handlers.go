package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Credentials - Create a struct to read the username and password from the request body
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

// TokenResponse -
type TokenResponse struct {
	Token string `json:"token"`
}

var jwtKey = []byte("my_secret_key")

// HandleServiceCheck - sends a simple message to verify that the service is up
func HandleServiceCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

// HandleAuth - handles auth requests
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var auth User

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil || auth.Password == "" || auth.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials format")
		return
	}

	user, err := auth.AuthenticateUser()

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Authentication failed")
		return
	}

	token, err := user.createTokenString()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusAccepted, &TokenResponse{Token: token})
}

// HandleRegister - handles auth registration of a new user
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register ")
}

// HandleCreateManager - handles auth registration of a new user
func HandleCreateManager(w http.ResponseWriter, r *http.Request) {
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
