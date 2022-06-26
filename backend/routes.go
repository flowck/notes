package main

import (
	"github.com/go-chi/chi/v5"
	"notes/modules/entry"
)

func RegisterRoutes(router chi.Router) {
	// Users

	// Entries
	router.Route("/entries", func(r chi.Router) {
		r.Get("/", entry.GetEntries)
		r.Post("/", entry.CreateEntry)
		r.Put("/{entryId}", entry.UpdateEntry)
	})

	// Folders
}
