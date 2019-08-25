package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

var jwtKey = []byte(os.Getenv("MOTUS_JWT_SECRET"))

func handleServiceCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "Service is running"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
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

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register ")
}

func (user *User) createTokenString() (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claimsExpiration := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	claims := &Claims{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Roles:          strings.Split(user.Roles, ","),
		StandardClaims: claimsExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
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
