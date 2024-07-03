package main

import (
	"fmt"
	"net/http"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	mux := http.NewServeMux()
	configs, err := cfg.Load()
	if err != nil {
		panic(err)
	}
	db, err := store.New(configs.Postgres)
	if err != nil {
		panic(err)
	}

	SetupRoutes(mux, db)

	fmt.Println("Starting MBOT service")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
