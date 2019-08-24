package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// InitServer - creates a new server with mux router
func InitServer() {
	fmt.Println("Server listening on port 9000. Started at: " + time.Now().Format(time.RFC3339))

	err := http.ListenAndServe(":9000", initRouter())

	if err != nil {
		log.Fatal("Failed to start server")
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleServiceCheck).Methods("GET")
	router.HandleFunc("/auth", HandleAuth).Methods("POST")
	router.HandleFunc("/register", HandleRegister).Methods("POST")

	return router
}
