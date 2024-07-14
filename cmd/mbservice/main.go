package main

import (
	"net/http"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/l"
)

var L *l.MLogger

func main() {
	cleanup := l.NewLogger()
	defer cleanup()

	mux := http.NewServeMux()
	configs, err := cfg.Load()
	if err != nil {
		panic(err)
	}

	SetupRoutes(mux, configs)

	l.Log("mbservice started")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
