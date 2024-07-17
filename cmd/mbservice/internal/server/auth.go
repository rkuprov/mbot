package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/l"
)

func (m *MBot) Login(ctx context.Context, req *connect.Request[mbotpb.LoginRequest]) (
	*connect.Response[mbotpb.LoginResponse],
	error,
) {
	token, err := m.auth.Login(ctx, req.Msg.GetUsername(), req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}

	resp := connect.Response[mbotpb.LoginResponse]{
		Msg: &mbotpb.LoginResponse{
			Ok: true,
		},
	}
	resp.Header().Set(auth.HeaderSessionToken, token)

	l.Log(fmt.Sprintf("User %s logged in", req.Msg.GetUsername()))

	return &resp, nil
}

func (m *MBot) Logout(ctx context.Context, req *connect.Request[mbotpb.LogoutRequest]) (
	*connect.Response[mbotpb.LogoutResponse],
	error,
) {
	err := m.auth.Logout(ctx, auth.SessionToken{
		Token: req.Header().Get(auth.HeaderSessionToken),
	})
	if err != nil {
		return nil, err
	}

	l.Log(fmt.Sprint("User logged out"))

	return &connect.Response[mbotpb.LogoutResponse]{}, nil
}

func (m *MBot) Register(ctx context.Context, req *connect.Request[mbotpb.RegisterRequest]) (
	*connect.Response[mbotpb.RegisterResponse],
	error,
) {
	err := m.auth.NewUser(ctx, req.Msg.GetUsername(), req.Msg.GetPassword())
	if err != nil {
		return nil, err
	}

	resp := connect.Response[mbotpb.RegisterResponse]{
		Msg: &mbotpb.RegisterResponse{
			Ok: true,
		},
	}

	l.Log(fmt.Sprintf("User %s registered", req.Msg.GetUsername()))

	return &resp, nil
}
