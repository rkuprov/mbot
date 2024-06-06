package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rkuprov/mbot/pkg/store"
)

func Status(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	}
}

func Login(ctx context.Context, db *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "token")

		active, date := db.ConfirmToken(ctx, id)
		if !active {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		_, err := fmt.Fprintf(w, "%s is valid until %s", id, date)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
