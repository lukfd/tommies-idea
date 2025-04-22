package main

import (
	"log"
	"net/http"

	"github.com/lukfd/redis-demo/model"
	"github.com/lukfd/redis-demo/routes"
)

func main() {
	redis := model.NewRedisClient()
	log.Println("Initializing Redis")
	redis.Init()

	// Initialize the default ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /store/{key}", func(w http.ResponseWriter, r *http.Request) {
		routes.Keyval(w, r, redis)
	})
	mux.HandleFunc("GET /idea/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.GetIdea(w, r, redis)
	})
	mux.HandleFunc("GET /ideas", func(w http.ResponseWriter, r *http.Request) {
		routes.GetIdeas(w, r, redis)
	})
	mux.HandleFunc("POST /idea", func(w http.ResponseWriter, r *http.Request) {
		routes.NewIdea(w, r, redis)
	})

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
