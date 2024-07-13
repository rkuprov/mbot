package server

import (
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/errs"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/store"
)

func (m *MBot) CreateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error) {
	id, err := m.db.CreateCustomer(ctx,
		store.CustomerCreate{
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"customers_email_key\"") {
			return nil, errs.HandleServerError(connect.CodeAlreadyExists, fmt.Errorf("customer with email %s already exists", req.Msg.GetEmail()))
		}
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	c, err := m.db.GetCustomer(ctx, id)
	if err != nil {
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	return &connect.Response[mbotpb.CreateCustomerResponse]{
		Msg: &mbotpb.CreateCustomerResponse{
			Message:         fmt.Sprintf("Customer created with ID: %s", id),
			Id:              c.GetId(),
			SubscriptionIds: c.GetSubscriptionIds(),
		},
	}, nil
}

func (m *MBot) GetCustomersAll(ctx context.Context,
	req *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error) {
	customers, err := m.db.GetCustomersAll(ctx)
	if err != nil {
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	out := make([]*mbotpb.Customer, 0)
	for _, c := range customers {
		out = append(out, &c)

	}
	return &connect.Response[mbotpb.GetCustomersAllResponse]{
		Msg: &mbotpb.GetCustomersAllResponse{
			Customers: out,
		},
	}, nil
}

func (m *MBot) GetCustomer(ctx context.Context,
	req *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error) {
	cst, err := m.db.GetCustomer(ctx, req.Msg.GetCustomerId())
	if err != nil {
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	if cst == nil {
		return nil, errs.HandleServerError(
			connect.CodeNotFound,
			fmt.Errorf("customer with ID %s not found", req.Msg.GetCustomerId()))
	}
	return &connect.Response[mbotpb.GetCustomerResponse]{
		Msg: &mbotpb.GetCustomerResponse{
			Customer: &mbotpb.Customer{
				Id:              cst.Id,
				Name:            cst.Name,
				Email:           cst.Email,
				Contact:         cst.Contact,
				SubscriptionIds: cst.SubscriptionIds,
			},
		},
	}, nil
}

func (m *MBot) UpdateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error) {
	err := m.db.UpdateCustomer(ctx, req.Msg.GetId(),
		store.CustomerUpdate{
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	out, err := m.db.GetCustomer(ctx, req.Msg.GetId())
	if err != nil {
		return nil, errs.HandleServerError(connect.CodeInternal, err)
	}
	return &connect.Response[mbotpb.UpdateCustomerResponse]{
		Msg: &mbotpb.UpdateCustomerResponse{
			Message:  fmt.Sprintf("Customer updated successfully"),
			Customer: out,
		},
	}, nil

}

func (m *MBot) DeleteCustomer(ctx context.Context,
	req *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error) {
	err := m.db.DeleteCustomer(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.DeleteCustomerResponse]{
		Msg: &mbotpb.DeleteCustomerResponse{
			Message: true,
		},
	}, nil
}
