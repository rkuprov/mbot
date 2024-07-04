package delete

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/cli/internal/ui"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Subscription struct {
	ID string `arg:"" help:"The ID of the subscription to delete" optional:""`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	_, err := client.DeleteSubscription(ctx, &connect.Request[mbotpb.DeleteSubscriptionRequest]{
		Msg: &mbotpb.DeleteSubscriptionRequest{
			SubscriptionId: s.ID,
		},
	})
	if err != nil {
		return err
	}

	ui.Tabular(ui.PrintCfg{
		Title: "Success!",
		Body:  []table.Row{{fmt.Sprintf("Subscription %s deleted successfully", s.ID)}},
	})

	return nil
}
