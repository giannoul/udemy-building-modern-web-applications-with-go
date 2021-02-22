package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/handlers"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create the template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
