package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/store"
)

func (m *MBot) CreateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error) {
	slug, err := m.db.CreateCustomer(ctx,
		store.CustomerCreate{
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		return nil, err
	}
	c, err := m.db.GetCustomer(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.CreateCustomerResponse]{
		Msg: &mbotpb.CreateCustomerResponse{
			Message:         fmt.Sprintf("Customer created with ID: %s", slug),
			Id:              c.GetId(),
			SubscriptionIds: c.GetSubscriptionIds(),
		},
	}, nil
}

func (m *MBot) GetCustomersAll(ctx context.Context,
	req *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error) {
	customers, err := m.db.GetCustomersAll(ctx)
	if err != nil {
		return nil, err
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
	cust, err := m.db.GetCustomer(ctx, req.Msg.GetCustomerId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.GetCustomerResponse]{
		Msg: &mbotpb.GetCustomerResponse{
			Customer: &mbotpb.Customer{
				Id:              cust.Id,
				Name:            cust.Name,
				Email:           cust.Email,
				Contact:         cust.Contact,
				SubscriptionIds: cust.SubscriptionIds,
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
		return nil, err
	}
	out, err := m.db.GetCustomer(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.UpdateCustomerResponse]{
		Msg: &mbotpb.UpdateCustomerResponse{
			Message:  fmt.Sprintf("Customer updated successfully"),
			Customer: &out,
		},
	}, nil

}

func (m *MBot) DeleteCustomer(ctx context.Context,
	req *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error) {
	// todo: implement delete customer
	// err := m.db.DeleteCustomer(ctx, req.Msg.GetSlug())
	// if err != nil {
	// 	return nil, err
	// }
	return &connect.Response[mbotpb.DeleteCustomerResponse]{
		Msg: &mbotpb.DeleteCustomerResponse{
			Message: fmt.Sprintf("Customer deleted with ID: %s", req.Msg.GetSlug()),
		},
	}, nil
}
