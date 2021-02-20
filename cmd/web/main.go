package main

import (
	"fmt"
	"net/http"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
