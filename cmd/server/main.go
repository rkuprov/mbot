package main

import (
	"context"
	"net/http"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	ctx := context.Background()
	r := http.NewServeMux()

	db, cancel := store.NewClient(ctx)
	m := &mServer{
		db: db,
	}
	defer cancel()

	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	r.Handle(path, handler)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
