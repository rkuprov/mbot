package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

type mServer struct {
	mbotpbconnect.UnimplementedMBotServerServiceHandler
	db *store.Store
}

func (m *mServer) CreateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error) {
	slug, err := m.db.CreateCustomer(ctx,
		store.CustomerCreate{
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		return nil, err
	}
	c, err := m.db.GetCustomer(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.CreateCustomerResponse]{
		Msg: &mbotpb.CreateCustomerResponse{
			Message:         fmt.Sprintf("Customer created with ID: %s", slug),
			Id:              c.GetId(),
			SubscriptionIds: c.GetSubscriptionIds(),
		},
	}, nil
}

func (m *mServer) GetCustomersAll(ctx context.Context,
	req *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error) {
	customers, err := m.db.GetCustomersAll(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*mbotpb.Customer, 0)
	for _, c := range customers {
		out = append(out, &c)

	}
	return &connect.Response[mbotpb.GetCustomersAllResponse]{
		Msg: &mbotpb.GetCustomersAllResponse{
			Customers: out,
		},
	}, nil
}

func (m *mServer) GetCustomer(ctx context.Context,
	req *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error) {
	cust, err := m.db.GetCustomer(ctx, req.Msg.GetSlug())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.GetCustomerResponse]{
		Msg: &mbotpb.GetCustomerResponse{
			Customer: &mbotpb.Customer{
				Id:              cust.Id,
				Name:            cust.Name,
				Email:           cust.Email,
				Contact:         cust.Contact,
				SubscriptionIds: cust.SubscriptionIds,
			},
		},
	}, nil
}

func (m *mServer) UpdateCustomer(ctx context.Context,
	req *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error) {
	err := m.db.UpdateCustomer(ctx, req.Msg.GetId(),
		store.CustomerUpdate{
			Name:    req.Msg.GetName(),
			Email:   req.Msg.GetEmail(),
			Contact: req.Msg.GetContact(),
		},
	)
	if err != nil {
		return nil, err
	}
	out, err := m.db.GetCustomer(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.UpdateCustomerResponse]{
		Msg: &mbotpb.UpdateCustomerResponse{
			Message:  fmt.Sprintf("Customer updated successfully"),
			Customer: &out,
		},
	}, nil

}

func (m *mServer) DeleteCustomer(ctx context.Context,
	req *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error) {
	err := m.db.DeleteCustomer(ctx, req.Msg.GetSlug())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.DeleteCustomerResponse]{
		Msg: &mbotpb.DeleteCustomerResponse{
			Message: fmt.Sprintf("Customer deleted with ID: %s", req.Msg.GetSlug()),
		},
	}, nil
}

func (m *mServer) CreateSubscription(ctx context.Context,
	req *connect.Request[mbotpb.CreateSubscriptionRequest]) (*connect.Response[mbotpb.CreateSubscriptionResponse], error) {
	start := req.Msg.GetSubscriptionStartDate().AsTime().Format("2006-01-02")
	id, err := m.db.CreateSubscription(ctx)
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
			Slug:         req.Msg.GetSlug(),
			Subscription: sub,
		},
	}, nil
}

func (m *mServer) GetSubscriptionsAll(ctx context.Context) (*connect.Response[mbotpb.GetSubscriptionsAllResponse], error) {
	// subs, err := m.db.GetSubscriptionsAll(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// custs, err := m.db.GetCustomersAll(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// out := make([]*mbotpb.GetSubscriptionResponse, 0)
	// for _, s := range subs {
	// 	for _, c := range custs {
	// 		if c.GetSlug() == s.GetSlug() {
	// 			cust = c.GetName()
	// 			break
	// 		}
	// 	}
	// 	out = append(out, &mbotpb.GetSubscriptionResponse{
	// 		Slug:            s.GetSlug(),
	// 		SubscriptionIds: s.GetSubscriptionIds(),
	// 		Customer:        cust,
	// 	})
	// }

	return &connect.Response[mbotpb.GetSubscriptionsAllResponse]{
		Msg: &mbotpb.GetSubscriptionsAllResponse{
			// Subscriptions: out,
		},
	}, nil
}

//
// func (m *mServer) UpdateSubscription(ctx context.Context,
// 	req *connect.Request[mbotpb.UpdateSubscriptionRequest]) (*connect.Response[mbotpb.UpdateSubscriptionResponse], error) {
// 	err := m.db.UpdateSubscription(ctx,
// 		req.Msg.GetId(),
// 		req.Msg.GetSlug(),
// 		req.Msg.GetSubscriptionExpiry(),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &connect.Response[mbotpb.UpdateSubscriptionResponse]{
// 		Msg: &mbotpb.UpdateSubscriptionResponse{
// 			Message: fmt.Sprintf("Subscription updated with ID: %s", req.Msg.GetId()),
// 		},
// 	}, nil
// }
//
// func (m *mServer) DeleteSubscription(ctx context.Context,
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
// func (m *mServer) GetSubcriptionByCustomer(ctx context.Context,
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
//
// func (m *mServer) GetStatsByCustomer(ctx context.Context,
// 	req *connect.Request[mbotpb.GetStatsByCustomerRequest]) (*connect.Response[mbotpb.GetStatsByCustomerResponse], error) {
// 	stats, err := m.db.GetStatsByCustomer(ctx, req.Msg.GetSlug())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &connect.Response[mbotpb.GetStatsByCustomerResponse]{
// 		Msg: &mbotpb.GetStatsByCustomerResponse{
// 			Stats: &mbotpb.Stats{
// 				Total: stats.Total,
// 				Used:  stats.Used,
// 			},
// 		},
// 	}, nil
// }
//
// func (m *mServer) GetStatsBySubscription(ctx context.Context,
// 	req *connect.Request[mbotpb.GetStatsBySubscriptionRequest]) (*connect.Response[mbotpb.GetStatsBySubscriptionResponse], error) {
// 	stats, err := m.db.GetStatsBySubscription(ctx, req.Msg.GetId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &connect.Response[mbotpb.GetStatsBySubscriptionResponse]{
// 		Msg: &mbotpb.GetStatsBySubscriptionResponse{
// 			Stats: &mbotpb.Stats{
// 				Total: stats.Total,
// 				Used:  stats.Used,
// 			},
// 		},
// 	}, nil
// }
//
// func (m *mServer) GetStatsAll(ctx context.Context) (*connect.Response[mbotpb.GetStatsAllResponse], error) {
// 	stats, err := m.db.GetStatsAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	out := make([]*mbotpb.Stats, 0)
// 	for _, s := range stats {
// 		out = append(out, &mbotpb.Stats{
// 			Total: s.Total,
// 			Used:  s.Used,
// 		})
// 	}
// 	return &connect.Response[mbotpb.GetStatsAllResponse]{
// 		Msg: &mbotpb.GetStatsAllResponse{
// 			Stats: out,
// 		},
// 	}, nil
// }
