package http

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"notes/internal/user/adapter"
	userService "notes/internal/user/service"
)

func InitUserHttpRoutes(router chi.Router, dbClient *sql.DB) {
	repository := adapter.NewUserPsqlRepository(dbClient)
	service := userService.NewUserService(repository)
	handlers := NewUserHttpHandler(&service)

	// Users
	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", handlers.Signup)
		r.Post("/signin", handlers.SignIn)
	})
}
