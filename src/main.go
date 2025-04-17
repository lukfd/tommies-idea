package main

import (
	"log"
	"net/http"

	"github.com/lukfd/redis-demo/internal"
	"github.com/lukfd/redis-demo/routes"
)

func main() {
	redis := internal.NewRedisClient()
	log.Println("Initializing Redis")
	redis.Init()

	// Initialize the default ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("GET /", routes.Index)
	mux.HandleFunc("GET /store/{key}", func(w http.ResponseWriter, r *http.Request) {
		routes.Keyval(w, r, redis)
	})
	mux.HandleFunc("GET /idea/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.Idea(w, r, redis)
	})

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
