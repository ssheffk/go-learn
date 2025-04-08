package main

import api "github.com/go-learn/go-middleware/middleware"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
