package main

import (
	"context"

	"github.com/go-chi/chi/v5"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/store"
)

func SetupRoutes(r *chi.Mux, db *store.Store) {
	ctx := context.Background()
	m := server.NewMBot(db)

	r.Get("/status", handlers.Status(ctx))

	// grpc-connect
	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	r.Handle(path, handler)

	r.Post("/subscription/{token}", handlers.Confirm(ctx, db))
}
