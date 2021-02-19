package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home page handler
func Home(respWritter http.ResponseWriter, req *http.Request) {
	n, err := fmt.Fprintf(respWritter, "This is the home page")
	fmt.Println(fmt.Sprintf("Response bytes count: %d", n))
	if err != nil {
		fmt.Println(err)
	}
}

// About page handler
func About(respWritter http.ResponseWriter, req *http.Request) {
	sum := addValue(2, 8)
	n, err := fmt.Fprintf(respWritter, fmt.Sprintf("This is the about page and 2+8=%d", sum))
	fmt.Println(fmt.Sprintf("Response bytes count: %d", n))
	if err != nil {
		fmt.Println(err)
	}
}

// Divide page handler
func Divide(respWritter http.ResponseWriter, req *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(respWritter, "Error, cannot divide by 0")
		return
	}
	fmt.Fprintf(respWritter, fmt.Sprintf("%f/%f=%f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// Just an addValue function
func addValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
