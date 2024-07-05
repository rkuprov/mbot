package main

import (
	"context"
	"errors"
	"net/http"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/store"
)

func SetupRoutes(mux *http.ServeMux, configs *cfg.Cfg) {
	ctx := context.Background()

	db, err := store.New(configs.Postgres)
	if err != nil {
		panic(err)
	}
	a, err := auth.New(configs.Postgres)
	if err != nil {
		panic(err)
	}
	m := server.NewMBot(db, a)

	// grpc-connect
	mux.Handle(mbotpbconnect.NewMBotServerServiceHandler(m, connect.WithInterceptors(WithTokenInterceptor(a))))

	mux.Handle("GET /status", handlers.Status(ctx))
	mux.Handle("POST /subscription/{token}", handlers.Confirm(ctx, db))
}

func WithTokenInterceptor(a *auth.Auth) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			id := req.Header().Get("id")
			if id == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("server malformed token"))
			}
			token := req.Header().Get("token")
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("server no token provided"))
			}
			newToken, err := a.ConfirmAndRotateToken(ctx, auth.SessionToken{
				UserID: id,
				Token:  token,
			})
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			resp, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			resp.Header().Set("token", newToken.Token)
			resp.Header().Set("id", newToken.UserID)

			return resp, nil
		}
	}
}
