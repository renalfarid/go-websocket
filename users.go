package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type usersResource struct{}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List) // GET /users - read a list of users

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list of stuff.."))
}
