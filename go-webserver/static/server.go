package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my simple Go web server!")
}

// handler for the hello page with a query parameter
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// handler for serving a static file
func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	// Route for the home page
	http.HandleFunc("/", homeHandler)

	// Route for the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Route for serving static files (e.g., HTML)
	http.HandleFunc("/static", staticFileHandler)

	// Start the server on port 8080
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
