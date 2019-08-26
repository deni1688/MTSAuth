package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register")
}

func (u *User) authenticate() (*User, error) {
	user, err := u.getUser(&User{Email: u.Email})

	if err != nil {
		return nil, err
	}

	if isValid := comparePasswords(user.Password, []byte(u.Password)); isValid {
		return user, nil
	}

	return nil, errors.New("passwordInvalid")
}

func (u *User) createTokenString() (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claimsExpiration := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	claims := &Claims{
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		Roles:          strings.Split(u.Roles, ","),
		StandardClaims: claimsExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		log.Println(err)

		return false
	}

	return true
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
