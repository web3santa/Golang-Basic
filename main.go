package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":3000"

// home is the ome page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THis is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is ather about apge and 2+ 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Connot divde by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y < 0 {
		err := errors.New("Connot divde by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// advalue add two integres and return the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
