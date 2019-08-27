package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	initDB()
	initServer()
}

func initDB() {
	db := connectDB()

	defer db.Close()

	db.AutoMigrate(&User{})
}

func initServer() {
	fmt.Println("Server listening on port 9000. Started at: " + time.Now().Format(time.RFC3339))
	router := initRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	if err := http.ListenAndServe(":9000", loggedRouter); err != nil {
		log.Fatal("Failed to start server")
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handleServiceCheck).Methods("GET")
	router.HandleFunc("/auth", handleLogin).Methods("POST")
	router.HandleFunc("/register", handleSignUp).Methods("POST")

	return router
}
