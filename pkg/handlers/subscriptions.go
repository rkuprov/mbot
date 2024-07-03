package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/rkuprov/mbot/pkg/store"
)

func Confirm(ctx context.Context, db *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if token == "" {
			http.Error(w, "token is required", http.StatusBadRequest)
			return
		}

		if err := db.ConfirmSubscription(ctx, token); err != nil {
			handleStoreError(w, err)
			return
		}
		if err := db.RecordUsage(ctx, token); err != nil {
			handleStoreError(w, err)
			return
		}

		w.Write([]byte("subscription active"))
	}
}

func handleStoreError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, store.ErrSubscriptionNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, store.ErrSubscriptionExpired):
		http.Error(w, err.Error(), http.StatusForbidden)
	case errors.Is(err, store.ErrSubscriptionNotActive):
		http.Error(w, err.Error(), http.StatusForbidden)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
