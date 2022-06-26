package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"net/http"
	"notes/infra"
	"strconv"
	"time"
)

func GlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,HEAD,OPTIONS")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	infra.InitDatabase(infra.Cfg)

	r.Use(GlobalMiddleware)
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(int(infra.Cfg.MaxRequestsPerMinute), 1*time.Minute))

	RegisterRoutes(r)

	fmt.Println("Server is running at port", infra.Cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(int(infra.Cfg.Port))), r)

	if err != nil {
		panic(err)
	}
}
