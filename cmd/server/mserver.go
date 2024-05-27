package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	SetupRoutes(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
