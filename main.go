package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database.DatabaseConnection()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to api"))
	})

	//r.Mount("/users", usersResource{}.Routes())
	r.Mount("/users", usersResource{}.Routes())
	r.Mount("/schools", schoolsResource{}.Routes())

	// Start API server in a Goroutine
	go func() {
		fmt.Println("API server listening on :3000")
		if err := http.ListenAndServe(":3000", r); err != nil {
			log.Fatalf("API server error: %s", err)
		}
	}()

	// Start WebSocket server
	server := newServer()
	http.Handle("/schools", websocket.Handler(server.handleScools))

	fmt.Println("WebSocket server listening on :3030")
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatalf("WebSocket server error: %s", err)
	}

}
