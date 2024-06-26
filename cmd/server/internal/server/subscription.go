package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/store"
)

func (m *MBot) CreateSubscription(ctx context.Context,
	req *connect.Request[mbotpb.CreateSubscriptionRequest]) (*connect.Response[mbotpb.CreateSubscriptionResponse], error) {
	start := req.Msg.GetSubscriptionStartDate().AsTime().Format("2006-01-02")
	id, err := m.db.CreateSubscription(ctx, store.SubscriptionCreate{
		CustomerID: req.Msg.GetCustomerId(),
		StartDate:  start,
		Duration:   int(req.Msg.GetDuration()),
	})
	if err != nil {
		return nil, err
	}
	sub, err := m.db.GetSubscription(ctx, id)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.CreateSubscriptionResponse]{
		Msg: &mbotpb.CreateSubscriptionResponse{
			Message:      "subscription created successfully",
			CustomerId:   req.Msg.GetCustomerId(),
			Subscription: sub,
		},
	}, nil
}

func (m *MBot) GetSubscription(ctx context.Context,
	req *connect.Request[mbotpb.GetSubscriptionRequest]) (*connect.Response[mbotpb.GetSubscriptionResponse], error) {
	sub, err := m.db.GetSubscription(ctx, req.Msg.GetSubscriptionId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.GetSubscriptionResponse]{
		Msg: &mbotpb.GetSubscriptionResponse{
			Subscription: sub,
		},
	}, nil
}

func (m *MBot) GetSubscriptionsAll(ctx context.Context, _ *connect.Request[mbotpb.GetSubscriptionsAllRequest]) (*connect.Response[mbotpb.GetSubscriptionsAllResponse], error) {
	subs, err := m.db.GetSubscriptionsAll(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*mbotpb.Subscription, 0)
	for _, s := range subs {
		out = append(out, &mbotpb.Subscription{
			SubscriptionId:     s.SubscriptionId,
			SubscriptionExpiry: s.SubscriptionExpiry,
		})
	}
	return &connect.Response[mbotpb.GetSubscriptionsAllResponse]{
		Msg: &mbotpb.GetSubscriptionsAllResponse{
			Subscriptions: out,
		},
	}, nil
}

func (m *MBot) GetSubscriptionByCustomer(ctx context.Context,
	req *connect.Request[mbotpb.GetSubscriptionByCustomerRequest]) (*connect.Response[mbotpb.GetSubscriptionByCustomerResponse], error) {
	sub, err := m.db.GetSubscriptionByCustomer(ctx, req.Msg.GetCustomerId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.GetSubscriptionByCustomerResponse]{
		Msg: &mbotpb.GetSubscriptionByCustomerResponse{
			Subscriptions: sub,
		},
	}, nil
}

func (m *MBot) UpdateSubscription(ctx context.Context,
	req *connect.Request[mbotpb.UpdateSubscriptionRequest]) (*connect.Response[mbotpb.UpdateSubscriptionResponse], error) {
	err := m.db.UpdateSubscription(ctx, store.SubscriptionUpdate{
		SubscriptionID: req.Msg.Id,
		StartDate:      req.Msg.StartDate,
		ExpirationDate: req.Msg.ExpirationDate,
	},
	)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.UpdateSubscriptionResponse]{
		Msg: &mbotpb.UpdateSubscriptionResponse{
			Message: fmt.Sprintf("Subscription %s has been updated.", req.Msg.Id),
		},
	}, nil
}

//
// func (m *MBot) DeleteSubscription(ctx context.Context,
// 	req *connect.Request[mbotpb.DeleteSubscriptionRequest]) (*connect.Response[mbotpb.DeleteSubscriptionResponse], error) {
// 	err := m.db.DeleteSubscription(ctx, req.Msg.GetId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &connect.Response[mbotpb.DeleteSubscriptionResponse]{
// 		Msg: &mbotpb.DeleteSubscriptionResponse{
// 			Message: fmt.Sprintf("Subscription deleted with ID: %s", req.Msg.GetId()),
// 		},
// 	}, nil
// }
//
// func (m *MBot) GetSubcriptionByCustomer(ctx context.Context,
// 	req *connect.Request[mbotpb.GetSubscriptionByCustomerRequest]) (*connect.Response[mbotpb.GetSubscriptionByCustomerResponse], error) {
// 	sub, err := m.db.GetSubscriptionByCustomer(ctx, req.Msg.GetSlug())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &connect.Response[mbotpb.GetSubscriptionByCustomerResponse]{
// 		Msg: &mbotpb.GetSubscriptionByCustomerResponse{
// 			Subscription: &mbotpb.Subscription{
// 				Id:                 sub.Id,
// 				Slug:               sub.Slug,
// 				SubscriptionExpiry: sub.SubscriptionExpiry,
// 			},
// 		},
// 	}, nil
// }
