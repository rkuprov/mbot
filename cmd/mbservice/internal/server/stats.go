package server

//
// func (m *MBot) GetStatsByCustomer(ctx context.Context,
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
// func (m *MBot) GetStatsBySubscription(ctx context.Context,
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
// func (m *MBot) GetStatsAll(ctx context.Context) (*connect.Response[mbotpb.GetStatsAllResponse], error) {
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
