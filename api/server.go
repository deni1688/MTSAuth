package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deni1688/motusauth/app"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Init ...
func Init(domain app.Domain) {
	fmt.Println("Server listening on port 9000. Started at: " + time.Now().Format(time.RFC3339))

	router := mux.NewRouter()

	c := &Controller{domain}

	router.HandleFunc("/", c.CheckServiceController).Methods("GET")
	router.HandleFunc("/login", c.LoginController).Methods("POST")
	router.HandleFunc("/register", c.RegisterController).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	if err := http.ListenAndServe(":9000", loggedRouter); err != nil {
		log.Fatal(err.Error())
	}
}
