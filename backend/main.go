package main

import (
	"log"
	"net/http"

	"github.com/noteapp/backend/pkg/handlers"
	"github.com/noteapp/backend/pkg/patterns"
	"github.com/noteapp/backend/pkg/storage"
)

func main() {
	// 1. Singleton Pattern: Get the single instance of our storage
	memStore := storage.GetMemoryStorage()

	// Initialize the handler with the storage (Strategy Pattern applied here as well)
	noteHandler := handlers.NewNoteHandler(memStore)

	// Apply Decorators (Middleware)
	decoratedHandler := patterns.ApplyDecorators(
		noteHandler.ServeHTTP,
		patterns.CORSDecorator,
		patterns.LoggingDecorator,
	)

	// Register the routes
	http.HandleFunc("/api/notes", decoratedHandler)
	http.HandleFunc("/api/notes/", decoratedHandler)

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
