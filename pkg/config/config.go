package config

import (
	"log"
	"text/template"
)

// AppConfig has the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
