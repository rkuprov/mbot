package update

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/cli/internal/ui"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Customer struct {
	ID      string `arg:"" required:"" help:"Customer ID"`
	Name    string `help:"Customer name"`
	Email   string `help:"Customer email"`
	Contact string `help:"Customer contact"`
}

func (c *Customer) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.UpdateCustomer(ctx, &connect.Request[mbotpb.UpdateCustomerRequest]{
		Msg: &mbotpb.UpdateCustomerRequest{
			Id:      c.ID,
			Name:    c.Name,
			Email:   c.Email,
			Contact: c.Contact,
		},
	})
	if err != nil {
		return err
	}

	var pc ui.PrintCfg
	switch {
	case resp.Msg == nil:
		pc.Title = "Failure!"
		pc.Body = []table.Row{{resp.Msg.GetMessage()}}
	default:
		pc.Title = fmt.Sprintf("Success! Updated customer ID: %s", resp.Msg.GetCustomer().GetId())
		pc.Header = table.Row{"ID", "Name", "Email", "Contact"}
		pc.Body = []table.Row{{
			resp.Msg.GetCustomer().GetId(),
			resp.Msg.GetCustomer().GetName(),
			resp.Msg.GetCustomer().GetEmail(),
			resp.Msg.GetCustomer().GetContact(),
		},
		}
	}

	ui.Tabular(pc)

	return nil
}
