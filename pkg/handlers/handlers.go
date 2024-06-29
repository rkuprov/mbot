package handlers

import (
	"context"
	"net/http"
)

func Status(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	}
}
