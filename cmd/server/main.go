package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	r := chi.NewRouter()
	configs := new(cfg.Cfg)
	err := configs.Load(filepath.Join("..", "..", "deployment", "config.json"))
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
