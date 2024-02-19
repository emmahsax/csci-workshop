package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIpToContext)
	mux.Use(app.Session.LoadAndSave)

	// Register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	mux.Route("/user", func(mux chi.Router) {
		mux.Use(app.auth)
		mux.Get("/profile", app.Profile)
	})

	// Static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("static", fileServer))

	return mux
}
