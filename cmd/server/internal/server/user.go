package server

import (
	"context"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

func (m *MBot) Login(ctx context.Context, req *connect.Request[mbotpb.LoginRequest]) (
	*connect.Response[mbotpb.LoginResponse],
	error,
) {
	token, err := m.auth.Login(ctx, req.Msg.GetUsername(), req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}
	return &connect.Response[mbotpb.LoginResponse]{
		Msg: &mbotpb.LoginResponse{
			Token: &mbotpb.SessionToken{
				UserId: token.UserID,
				Token:  token.Token,
			},
		},
	}, nil
}
