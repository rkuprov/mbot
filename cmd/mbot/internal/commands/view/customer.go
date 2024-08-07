package view

import (
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/mbot/internal/ui"
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
		if strings.Contains(err.Error(), pgx.ErrNoRows.Error()) {
			ui.Tabular(ui.PrintCfg{
				Body: []table.Row{{fmt.Sprintf("Customer with ID %s not found", id)}},
			})
			return nil
		}
		return err
	}

	var pc ui.PrintCfg
	switch {
	case resp.Msg == nil:
		pc.Title = "Failure!"
		pc.Body = []table.Row{{resp.Msg.String()}}
	default:
		pc.Title = fmt.Sprintf("Success!")
		subCount := len(resp.Msg.Customer.GetSubscriptionIds())
		pc.Header = table.Row{"ID", "Name", "Email", "Contact", fmt.Sprintf("Subscription ID (%d)", subCount)}
		row := table.Row{
			resp.Msg.Customer.GetId(),
			resp.Msg.Customer.GetName(),
			resp.Msg.Customer.GetEmail(),
			resp.Msg.Customer.GetContact(),
		}
		if len(resp.Msg.Customer.GetSubscriptionIds()) == 0 {
			pc.Body = []table.Row{row}
			break
		}
		for i, sub := range resp.Msg.Customer.GetSubscriptionIds() {
			if i == 0 {
				pc.Body = []table.Row{append(row, sub)}
				continue
			}
			pc.Body = append(pc.Body, table.Row{"", "", "", "", sub})
		}
	}

	ui.Tabular(pc)

	return nil
}

func viewAllCustomers(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.GetCustomersAll(ctx, &connect.Request[mbotpb.GetCustomersAllRequest]{})
	if err != nil {
		return err
	}

	var pc ui.PrintCfg
	switch {
	case resp == nil:
		pc.Title = "Failure!"
	case len(resp.Msg.GetCustomers()) == 0:
		pc.Body = []table.Row{{"No active customers found"}}
	default:
		pc.Title = fmt.Sprintf("Success! Found %d active customers", len(resp.Msg.GetCustomers()))
		pc.Header = table.Row{"ID", "Name", "Email", "Contact", "Subscription Count"}
		for _, cust := range resp.Msg.GetCustomers() {
			pc.Body = append(pc.Body, table.Row{
				cust.GetId(),
				cust.GetName(),
				cust.GetEmail(),
				cust.GetContact(),
				len(cust.GetSubscriptionIds()),
			})
		}

	}

	ui.Tabular(pc)

	return nil
}
