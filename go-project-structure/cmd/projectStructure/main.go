package main

import (
	"fmt"
	"net/http"

	"github.com/ssheffk/go-learn/go-project-structure/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
