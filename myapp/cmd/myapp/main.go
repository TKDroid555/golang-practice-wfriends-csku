package main

import (
	"log"
	"net/http"

	"myapp/internal/config"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/service"
)

func main() {
	cfg := config.Load()

	// Wire dependencies bottom-up: repository -> service -> handler.
	// This is plain constructor injection — the idiomatic way to do
	// DI in Go, no framework/container needed.
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()
	userHandler.Routes(mux)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := ":" + cfg.Port
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
