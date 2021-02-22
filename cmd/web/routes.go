package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// http.Handler is also known as multiplexer
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
