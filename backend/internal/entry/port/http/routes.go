package http

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	entryAdapter "notes/internal/entry/adapter"
	entryService "notes/internal/entry/service"
)

func InitEntryHttpRoutes(router chi.Router, dbClient *sql.DB) {
	repository := entryAdapter.NewEntryPsqlRepository(dbClient)
	service := entryService.NewEntryService(repository)
	handlers := NewEntryHttpHandler(&service)

	router.Route("/entries", func(r chi.Router) {
		r.Get("/", handlers.GetEntries)
		r.Post("/", handlers.CreateEntry)
		// r.Put("/{entryId}", handlers.UpdateEntry)
		r.Get("/{entryId}", handlers.GetEntry)
	})
}
