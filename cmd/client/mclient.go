package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

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

	resp2, err := c.CreateSubscription(context.Background(), &connect.Request[mbotpb.CreateSubscriptionRequest]{Msg: &mbotpb.CreateSubscriptionRequest{
		Slug:                  "john-doe",
		SubscriptionStartDate: timestamppb.Now(),
		Duration:              30,
	}})
	if err != nil {
		panic(err)
	}
	println(resp2.Msg.String())
}
