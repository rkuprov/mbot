package main

import (
	"context"
	"errors"
	"net/http"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/cmd/mbservice/internal/server"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/errs"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/handlers"
	"github.com/rkuprov/mbot/pkg/l"
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
	// server handler
	mux.Handle(mbotpbconnect.NewMBotServerServiceHandler(m, connect.WithInterceptors(WithTokenInterceptor(a))))
	// auth handler
	mux.Handle(mbotpbconnect.NewMbotAuthServerServiceHandler(m))

	mux.Handle("GET /status", handlers.Status(ctx))
	mux.Handle("POST /subscription/{token}", handlers.Confirm(ctx, db))
	mux.Handle("GET /subscription/{token}/version", handlers.Version(ctx))
	mux.Handle("GET /subscription/{token}/update", handlers.Update(ctx, db))
	l.Log("routes setup")
}

func WithTokenInterceptor(a *auth.Auth) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			token := req.Header().Get(auth.HeaderSessionToken)
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("mbservice no token provided"))
			}
			newToken, err := a.ConfirmAndRotateToken(ctx, token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			resp, err := next(ctx, req)
			if err != nil {
				return resp, errs.HandleServerError(connect.CodeInternal, err, errs.WithSessionTokenDetail(newToken.Value))
			}

			resp.Header().Set(auth.HeaderSessionToken, newToken.Value)

			return resp, nil
		}
	}
}
