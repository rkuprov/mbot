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
	EndDate   time.Time `format:"2006-01-02" aliases:"ed,e" default:"0001-01-01" help:"End date of the subscription (YYYY-MM-DD)"`
	Duration  *int      `aliases:"dur,d" help:"Duration of the subscription in days"`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	startDate, endDate, err := s.toSubscriptionDates()
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

	fmt.Println(resp.Msg.GetMessage())

	return nil
}

func (s *Subscription) toSubscriptionDates() (time.Time, time.Time, error) {
	var endDate time.Time
	switch {
	case s.Duration == nil && s.EndDate.IsZero():
		return time.Time{}, time.Time{}, fmt.Errorf("either duration or end date must be provided")
	case s.Duration == nil:
		endDate = s.EndDate
	default:
		endDate = s.StartDate.AddDate(0, 0, *s.Duration)
	}

	if s.StartDate.After(endDate) {
		return time.Time{}, time.Time{}, fmt.Errorf("start date cannot be after end date")
	}

	return s.StartDate, endDate, nil
}
