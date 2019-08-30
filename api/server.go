package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Init ...
func Init() {
	fmt.Println("Server listening on port 9000. Started at: " + time.Now().Format(time.RFC3339))
	router := initRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	if err := http.ListenAndServe(":9000", loggedRouter); err != nil {
		log.Fatal(err.Error())
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", CheckServiceController).Methods("GET")
	router.HandleFunc("/login", LoginController).Methods("POST")
	router.HandleFunc("/signup", SignUpController).Methods("POST")

	return router
}
