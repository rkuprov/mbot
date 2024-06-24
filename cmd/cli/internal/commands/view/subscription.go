package view

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Subscription struct {
	ID         string `arg:"" help:"The ID of the subscription to view" optional:""`
	All        bool   `help:"View all subscriptions"`
	CustomerID string `aliases:"c,cid,for" help:"View a subscription for a customer" optional:""`
}

func (s *Subscription) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	switch {
	case s.All:
		return viewAllSubscriptions(ctx, client)
	case s.CustomerID != "":
		return viewSubscritpionByCustomer(ctx, client, s.CustomerID)
	default:
		return viewSubscription(ctx, client, s.ID)
	}

}

func viewSubscription(ctx context.Context, client mbotpbconnect.MBotServerServiceClient, id string) error {
	resp, err := client.GetSubscription(ctx, &connect.Request[mbotpb.GetSubscriptionRequest]{
		Msg: &mbotpb.GetSubscriptionRequest{
			SubscriptionId: id,
		},
	})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg)

	return nil
}

func viewAllSubscriptions(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.GetSubscriptionsAll(ctx, &connect.Request[mbotpb.GetSubscriptionsAllRequest]{})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg)

	return nil
}

func viewSubscritpionByCustomer(ctx context.Context, client mbotpbconnect.MBotServerServiceClient, id string) error {
	resp, err := client.GetSubscriptionByCustomer(ctx, &connect.Request[mbotpb.GetSubscriptionByCustomerRequest]{
		Msg: &mbotpb.GetSubscriptionByCustomerRequest{
			CustomerId: id,
		},
	})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg)

	return nil
}
