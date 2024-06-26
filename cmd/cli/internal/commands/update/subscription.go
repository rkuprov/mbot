package update

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Subscription struct {
	ID             string    `arg:"" required:"" help:"Subscription ID"`
	StartDate      time.Time `optional:"" format:"2006-01-02" default:"0001-01-01"  help:"Start date"`
	ExpirationDate time.Time `optional:"" format:"2006-01-02" default:"0001-01-01"  help:"Expiration date"`
	Duration       *int      `optional:"" default:"0" help:"Duration of the subscription in days"`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	start := new(timestamppb.Timestamp)
	if !s.StartDate.IsZero() {
		start = timestamppb.New(s.StartDate)
	}
	expiration := new(timestamppb.Timestamp)
	if !s.ExpirationDate.IsZero() {
		expiration = timestamppb.New(s.ExpirationDate)
	}
	resp, err := client.UpdateSubscription(ctx, &connect.Request[mbotpb.UpdateSubscriptionRequest]{
		Msg: &mbotpb.UpdateSubscriptionRequest{
			Id:             s.ID,
			StartDate:      start,
			ExpirationDate: expiration,
		}})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg.GetMessage())

	return nil
}
