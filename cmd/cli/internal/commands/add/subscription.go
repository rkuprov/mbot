package add

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
	ID        string    `arg:"" help:"Customer ID" required:""`
	StartDate time.Time `format:"2006-01-02" aliases:"sd,s" required:"" help:"Start date of the subscription (YYYY-MM-DD)"`
	Duration  int       `required:"" aliases:"dur,d" help:"Duration of the subscription in days"`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.CreateSubscription(ctx, &connect.Request[mbotpb.CreateSubscriptionRequest]{Msg: &mbotpb.CreateSubscriptionRequest{
		CustomerId:            s.ID,
		SubscriptionStartDate: timestamppb.New(s.StartDate),
		Duration:              int32(s.Duration),
	},
	})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg.GetMessage())

	return nil
}
