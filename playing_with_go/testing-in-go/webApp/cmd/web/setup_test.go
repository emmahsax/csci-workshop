package main

import (
	"log"
	"os"
	"testing"
	"webApp/pkg/db"
)

// This file (must be named setup_test.go) and the TestMain() function will always be run
// at the beginning of running the tests. This allows us to set up a test environment to
// use in all tests.

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()
	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
