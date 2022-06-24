package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"net/http"
	"notes/internal/infra"
	"notes/internal/repositories"
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
	ctx, _ := context.WithCancel(context.Background())

	infra.InitDatabase(infra.Cfg)

	r.Use(GlobalMiddleware)
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(int(infra.Cfg.MaxRequestsPerMinute), 1*time.Minute))

	// entries, ok := repositories.FindEntries(ctx)
	// folders, ok := repositories.FindFolders(ctx)
	entry, ok := repositories.FindEntryById(ctx, "64086fd9-2eba-4e8d-972e-e79e10c74d42")

	if ok != nil {
		panic(ok)
	}

	// fmt.Println(entries)
	// fmt.Println(folders)
	fmt.Println(entry)

	fmt.Println("Server is running at port", infra.Cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(int(infra.Cfg.Port))), r)

	if err != nil {
		panic(err)
	}
}
