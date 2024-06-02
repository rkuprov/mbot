package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

func main() {
	ctx := context.Background()
	r := http.NewServeMux()

	s := grpc.NewServer()
	db, cancel := store.NewClient(ctx)
	m := &mServer{
		db: db,
	}
	defer cancel()

	path, handler := mbotpbconnect.NewMBotServerServiceHandler(m)
	r.Handle(path, handler)

	mbotpbconnect.MBotServerServiceClient(s, m)
	reflection.Register(s)
	customers, err := db.GetCustomersAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
