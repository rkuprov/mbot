package update

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

const timeFormat = "2006-01-02"

type Subscription struct {
	ID             string    `arg:"" required:"" help:"Subscription ID"`
	StartDate      time.Time `optional:"" format:"2006-01-02" default:"0001-01-01"  help:"Start date"`
	ExpirationDate time.Time `optional:"" format:"2006-01-02" default:"0001-01-01"  help:"Expiration date"`
	Duration       *int      `optional:"" default:"0" help:"Duration of the subscription in days"`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	start, expiration, err := commands.ToSubscriptionDates(s.StartDate, s.ExpirationDate, s.Duration)
	if err != nil {
		return err
	}

	resp, err := client.UpdateSubscription(ctx, &connect.Request[mbotpb.UpdateSubscriptionRequest]{
		Msg: &mbotpb.UpdateSubscriptionRequest{
			Id:             s.ID,
			StartDate:      timestamppb.New(start),
			ExpirationDate: timestamppb.New(expiration),
		}})
	if err != nil {
		return err
	}

	var pc ui.PrintCfg
	switch {
	case resp.Msg == nil:
		pc.Title = "Failure!"
		pc.Body = []table.Row{{resp.Msg.String()}}
	default:
		pc.Title = fmt.Sprintf("Success! Updated subscription ID: %s", resp.Msg.GetId())
		pc.Header = table.Row{
			"Status",
			"Start Date",
			"Expiration Date",
		}
		pc.Body = []table.Row{
			{"Original", resp.Msg.StartDate.AsTime().Format(timeFormat), resp.Msg.ExpirationDate.AsTime().Format(timeFormat)},
		}
		pc.Body = append(pc.Body, table.Row{
			"Updated", resp.Msg.UpdatedStartDate.AsTime().Format(timeFormat), resp.Msg.UpdatedExpirationDate.AsTime().Format(timeFormat)})

	}

	ui.Tabular(pc)

	return nil
}
