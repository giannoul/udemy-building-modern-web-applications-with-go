package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is just a custom middleware
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This came from middleware")
		next.ServeHTTP(w, r)
	})
}

// NoSurf middleware for CSRF
func NoSurf(next http.Handler) http.Handler {
	csrfHander := nosurf.New(next)
	csrfHander.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHander
}
