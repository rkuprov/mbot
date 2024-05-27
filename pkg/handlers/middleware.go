package handlers

import (
	"context"
	"mbot/pkg/store"
	"net/http"
)

func AuthMiddleware(ctx context.Context, db *store.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if !db.ConfirmToken(ctx, token) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			//TODO: rotate token after confirmation
			next.ServeHTTP(w, r)
		})
	}
}
