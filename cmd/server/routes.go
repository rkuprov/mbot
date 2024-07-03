package main

import (
	"context"
	"net/http"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/store"
)

func SetupRoutes(mux *http.ServeMux, db *store.Store) {
	ctx := context.Background()
	m := server.NewMBot(db)

	// grpc-connect
	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	mux.Handle(path, handler)

	mux.Handle("GET /status", handlers.Status(ctx))
	mux.Handle("POST /subscription/{token}", handlers.Confirm(ctx, db))
}
