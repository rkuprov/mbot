package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	r := chi.NewRouter()
	configs, err := cfg.Load()
	if err != nil {
		panic(err)
	}
	db, err := store.New(configs.Postgres)
	if err != nil {
		panic(err)
	}

	SetupRoutes(r, db)

	fmt.Println("Starting MBOT service")
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
