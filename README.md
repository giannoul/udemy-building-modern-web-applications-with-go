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
│       ├── main.go
│       ├── middleware.go
│       └── routes.go
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── handlers
│   │   └── handlers.go
│   ├── models
│   │   └── templatedata.go
│   └── render
│       └── render.go
├── README.md
└── templates
    ├── about.page.tmpl
    ├── base.layout.tmpl
    └── home.page.tmpl
```


## Go modules

In order to start using go modules, in the root directory issue the command:

```
go mod init github.com/giannoul/udemy-building-modern-web-applications-with-go
```