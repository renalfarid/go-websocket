package main

import (
	"log"
	"net/http"
	"rest-api/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
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

	http.ListenAndServe(":3000", r)

}
