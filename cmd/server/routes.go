package main

import (
	"context"

	"github.com/go-chi/chi/v5"

	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/store"
)

func SetupRoutes(r *chi.Mux) {
	ctx := context.Background()
	db := store.NewClient(ctx)

	r.Get("/status", handlers.Status(ctx))
	r.Post("/login/{token}", handlers.Login(ctx, db))
}
