package main

import (
	"fmt"
	"net/http"

	"github.com/rkuprov/mbot/pkg/cfg"
)

func main() {
	mux := http.NewServeMux()
	configs, err := cfg.Load()
	if err != nil {
		panic(err)
	}

	SetupRoutes(mux, configs)

	fmt.Println("Starting MBOT service")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
