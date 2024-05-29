package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

func main() {
	c := mbotpbconnect.NewMBotServerClient(http.DefaultClient, "http://localhost:8080")
	resp, err := c.CreateCustomer(context.Background(), &connect.Request[mbotpb.CreateCustomerRequest]{
		Msg: &mbotpb.CreateCustomerRequest{
			Id:      "1",
			Name:    "John Doe",
			Email:   "doe@gmail.com",
			Contact: "1234567890",
			Token:   uuid.New().String(),
		},
	})
	if err != nil {
		panic(err)
	}

	println(resp.Msg.String())
}
