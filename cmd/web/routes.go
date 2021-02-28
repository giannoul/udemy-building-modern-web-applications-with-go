package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// http.Handler is also known as multiplexer
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))*/

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(middleware.Logger)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
