package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

func main() {
	// r := chi.NewRouter()
	// SetupRoutes(r)
	r := http.NewServeMux()

	path, handler := mbotpbconnect.NewMBotServerHandler(&mbotserver{})
	r.Handle(path, handler)
	fmt.Println(path)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}

type mbotserver struct {
	mbotpbconnect.UnimplementedMBotServerHandler
}

func (m mbotserver) CreateCustomer(c context.Context,
	req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerReply], error) {
	return &connect.Response[mbotpb.CreateCustomerReply]{
		Msg: &mbotpb.CreateCustomerReply{
			Message: "success",
			Id:      req.Msg.GetId(),
			Token:   req.Msg.GetToken(),
		},
	}, nil
}
