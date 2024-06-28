package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	r := http.NewServeMux()
	configs := new(cfg.Cfg)
	err := configs.Load(filepath.Join("..", "..", "deployment", "config.json"))
	if err != nil {
		panic(err)
	}
	db, err := store.New(configs.Postgres)
	if err != nil {
		panic(err)
	}
	m := server.NewMBot(db)

	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	r.Handle(path, handler)

	fmt.Println("Starting MBOT service")
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
