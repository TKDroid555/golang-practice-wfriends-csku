package patterns

import (
	"log"
	"net/http"
	"time"
)

// Decorator Pattern: Wraps an http.HandlerFunc to add extra behavior
// (like logging and CORS) without modifying the original handler's code.

// Middleware is a type for our decorator functions
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyDecorators applies a list of decorators to a handler
func ApplyDecorators(h http.HandlerFunc, decorators ...Middleware) http.HandlerFunc {
	for _, decorator := range decorators {
		h = decorator(h)
	}
	return h
}

// LoggingDecorator logs the method, path, and duration of the request
func LoggingDecorator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		log.Printf("[%s] %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// CORSDecorator adds CORS headers to the response
func CORSDecorator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
