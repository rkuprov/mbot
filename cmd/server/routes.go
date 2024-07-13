package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/cmd/server/internal/server"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/errs"
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
	// server handler
	mux.Handle(mbotpbconnect.NewMBotServerServiceHandler(m, connect.WithInterceptors(WithTokenInterceptor(a))))
	// auth handler
	mux.Handle(mbotpbconnect.NewMbotAuthServerServiceHandler(m))

	mux.Handle("GET /status", handlers.Status(ctx))
	mux.Handle("POST /subscription/{token}", handlers.Confirm(ctx, db))
}

func WithTokenInterceptor(a *auth.Auth) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			token := req.Header().Get(auth.HeaderSessionToken)
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("server no token provided"))
			}
			newToken, err := a.ConfirmAndRotateToken(ctx, token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			resp, err := next(ctx, req)
			if err != nil {
				return resp, errs.HandleServerError(connect.CodeInternal, err, errs.WithSessionTokenDetail(newToken.Value))
			}

			fmt.Printf("new token: %s @ %s\n", newToken.Value, time.Now().String())
			resp.Header().Set(auth.HeaderSessionToken, newToken.Value)

			return resp, nil
		}
	}
}
