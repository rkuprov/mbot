package add

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/cmd/mbot/internal/commands"
	"github.com/rkuprov/mbot/cmd/mbot/internal/ui"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Subscription struct {
	ID        string    `arg:"" help:"Customer ID" required:""`
	StartDate time.Time `format:"2006-01-02" aliases:"sd,s" required:"" help:"Start date of the subscription (YYYY-MM-DD)"`
	EndDate   time.Time `format:"2006-01-02" aliases:"ed,e" default:"0001-01-01" help:"End date of the subscription (YYYY-MM-DD)"`
	Duration  *int      `aliases:"dur,d" help:"Duration of the subscription in days"`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	startDate, endDate, err := commands.ToSubscriptionDates(s.StartDate, s.EndDate, s.Duration)
	if err != nil {
		return err
	}
	resp, err := client.CreateSubscription(ctx, &connect.Request[mbotpb.CreateSubscriptionRequest]{Msg: &mbotpb.CreateSubscriptionRequest{
		CustomerId:     s.ID,
		StartDate:      timestamppb.New(startDate),
		ExpirationDate: timestamppb.New(endDate),
	},
	})
	if err != nil {
		return err
	}

	var pc ui.PrintCfg

	switch resp.Msg {
	case nil:
		pc.Title = "Failure!"
		pc.Body = []table.Row{{resp.Msg.GetMessage()}}
	default:
		pc.Title = fmt.Sprintf("Success! Subscription created for customer ID: %s", resp.Msg.GetSubscription().GetCustomerId())
		pc.Header = table.Row{"ID", "Start Date", "Expiration Date"}
		pc.Body = []table.Row{{resp.Msg.GetSubscription().SubscriptionId,
			resp.Msg.GetSubscription().GetStartDate().AsTime().Format("2006-01-02"),
			resp.Msg.GetSubscription().GetExpirationDate().AsTime().Format("2006-01-02")}}
	}

	ui.Tabular(pc)
	return nil
}
