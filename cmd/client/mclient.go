package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

func main() {
	c := mbotpbconnect.NewMBotServerServiceClient(http.DefaultClient, "http://localhost:8080")
	resp, err := c.CreateCustomer(context.Background(), &connect.Request[mbotpb.CreateCustomerRequest]{
		Msg: &mbotpb.CreateCustomerRequest{
			Name:    "John Doe",
			Email:   "doe@gmail.com",
			Contact: "1234567890",
		},
	})
	if err != nil {
		panic(err)
	}

	println(resp.Msg.String())
}
