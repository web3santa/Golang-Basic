package main

import (
	"fmt"
	"net/http"

	"github.com/santswap/go-basic/pkg/handlers"
)

const portNumber = ":3000"

// main is the main application
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
