package delete

import (
	"context"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/cli/internal/ui"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Customer struct {
	ID string `arg:"" help:"The ID of the customer to delete"`
}

func (c *Customer) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	_, err := client.DeleteCustomer(ctx, &connect.Request[mbotpb.DeleteCustomerRequest]{Msg: &mbotpb.DeleteCustomerRequest{Id: c.ID}})
	if err != nil {
		return err
	}

	ui.Tabular(ui.PrintCfg{
		Title: "Success!",
		Body:  []table.Row{{"Customer %s was deleted successfully", c.ID}},
	})
	return nil
}
