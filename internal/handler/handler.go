package handler

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rafaeldepontes/auth-go/api"
)

// Handler controls the system routes based on *chi.Mux and a configuration struct.
func Handler(r *chi.Mux, app *api.Application) {
	r.Use(chimiddleware.StripSlashes)

	// Public
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", app.UserService.FindAllUsers)
		// r.Get("/users/{id}", app.UserService.FindAllUsers)
	})

	// Protected
	// WIP
}
