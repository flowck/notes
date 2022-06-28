package main

import (
	"github.com/go-chi/chi/v5"
	"notes/modules/entry"
	"notes/modules/user"
)

func RegisterRoutes(router chi.Router) {
	// Users
	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", user.Signup)
		r.Post("/signin", user.SignIn)
	})

	// Entries
	router.Route("/entries", func(r chi.Router) {
		r.Get("/", entry.GetEntries)
		r.Post("/", entry.CreateEntry)
		r.Put("/{entryId}", entry.UpdateEntry)
		r.Get("/{entryId}", entry.GetEntry)
	})

	// Folders
}
