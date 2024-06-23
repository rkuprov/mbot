package view

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Customer struct {
	ID  string `arg:"" help:"The ID of the customer to view" optional:""`
	All bool   `help:"View all customers"`
}

func (customer *Customer) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	if customer.All {
		return viewAllCustomers(ctx, client)
	}

	return viewCustomer(ctx, client, customer.ID)
}

func viewCustomer(ctx context.Context, client mbotpbconnect.MBotServerServiceClient, id string) error {
	resp, err := client.GetCustomer(ctx, &connect.Request[mbotpb.GetCustomerRequest]{
		Msg: &mbotpb.GetCustomerRequest{
			CustomerId: id,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(resp.Msg)

	return nil
}

func viewAllCustomers(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.GetCustomersAll(ctx, &connect.Request[mbotpb.GetCustomersAllRequest]{})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg)

	return nil
}
