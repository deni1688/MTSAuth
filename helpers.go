package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	if err != nil {
		return ""
	}

	return string(hash)
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
