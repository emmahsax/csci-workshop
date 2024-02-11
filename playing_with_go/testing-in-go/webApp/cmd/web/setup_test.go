package main

import (
	"os"
	"testing"
)

// This file (must be named setup_test.go) and the TestMain() function will always be run
// at the beginning of running the tests. This allows us to set up a test environment to
// use in all tests.

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()

	os.Exit(m.Run())
}
