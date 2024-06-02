package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

type mServer struct {
	mbotpbconnect.UnimplementedMBotServerHandler
	db *store.Client
}

func (m mServer) CreateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerReply], error) {
	slug, err := m.db.CreateCustomer(ctx,
		req.Msg.GetName(),
		req.Msg.GetEmail(),
		req.Msg.GetContact(),
	)
	if err != nil {
		return nil, err
	}
	cust, err := m.db.GetCustomer(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.CreateCustomerReply]{
		Msg: &mbotpb.CreateCustomerReply{
			Message:        fmt.Sprintf("Customer created with ID: %s", slug),
			Slug:           cust.Slug,
			SubscriptionId: cust.SubscriptionId,
		},
	}, nil
}
