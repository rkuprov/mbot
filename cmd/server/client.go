package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	"github.com/rkuprov/mbot/pkg/datamodel"
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

	id, err := m.db.CreateCustomer(ctx,
		datamodel.Customer{
			ID:      uuid.New().String(),
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.CreateCustomerReply]{
		Msg: &mbotpb.CreateCustomerReply{
			Message: fmt.Sprintf("Customer created with ID: %s", id),
			Id:      req.Msg.GetId(),
			Token:   req.Msg.GetToken(),
		},
	}, nil
}
