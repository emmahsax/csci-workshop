package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// Set up an app config
	app := application{}

	// Get application routes
	mux := app.routes()

	// Get a session manager
	app.Session = getSession()

	// Print out a message
	log.Println("Starting server on port 8080... go to http://localhost:8080/")

	// Start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
