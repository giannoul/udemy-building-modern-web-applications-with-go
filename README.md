# udemy-building-modern-web-applications-with-go

## This repository has the project's code for the course

## How to run the application

```
go run cmd/web/*.go
```


## Directory structure

It is common to put the `main` package in `cmd/web`. At this moment the directory structure is 

```
.
├── cmd
│   └── web
│       └── main.go #the main package + entry point of the application
├── go.mod
├── pkg
│   ├── handlers
│   │   └── handlers.go #the route hanlders
│   └── render
│       └── render.go   #the template renderer
├── README.md
└── templates
    ├── about.page.tmpl #page template
    └── home.page.tmpl  # ^ same thing
```


## Go modules

In order to start using go modules, in the root directory issue the command:

```
go mod init github.com/giannoul/udemy-building-modern-web-applications-with-go
```