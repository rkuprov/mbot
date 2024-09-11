package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/rkuprov/mbot/pkg/store"
)

func Update(ctx context.Context, db *store.Store) http.HandlerFunc {
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

		files, err := os.ReadDir(filepath.Join("..", "data"))
		if err != nil {
			handleStoreError(w, err)
			return
		}

		versions := make([]string, 0, len(files))
		for _, file := range files {
			versions = append(versions, file.Name())
		}
		slices.Sort(versions)
		if len(versions) == 0 {
			http.Error(w, "no versions available", http.StatusNotFound)
			return
		}

		toUpdate, err := os.ReadFile(filepath.Join("..", "data", versions[len(versions)-1]))
		if err != nil {
			handleStoreError(w, err)
			return
		}
		checksum := fmt.Sprintf("%x", sha256.Sum256(toUpdate))

		w.Header().Set("mbot-checksum", checksum)
		w.Write(toUpdate)
	}
}

func Version(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		version, err := os.ReadFile(filepath.Join("..", "VERSION"))
		if err != nil {
			handleStoreError(w, err)
			return
		}

		w.Write([]byte(version))
	}
}
