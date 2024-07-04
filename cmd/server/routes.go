package main

import (
	"context"
	"net/http"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/store"
)

func SetupRoutes(mux *http.ServeMux, configs *cfg.Cfg) {
	ctx := context.Background()

	db, err := store.New(configs.Postgres)
	if err != nil {
		panic(err)
	}
	a, err := auth.New(configs.Postgres)
	if err != nil {
		panic(err)
	}
	m := server.NewMBot(db, a)

	// grpc-connect
	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	mux.Handle(path, handler)

	mux.Handle("GET /status", handlers.Status(ctx))
	mux.Handle("POST /subscription/{token}", handlers.Confirm(ctx, db))
}
