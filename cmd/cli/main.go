package main

import (
	"context"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/alecthomas/kong"

	"github.com/rkuprov/mbot/cmd/cli/internal/commands/add"
	"github.com/rkuprov/mbot/cmd/cli/internal/commands/auth"
	"github.com/rkuprov/mbot/cmd/cli/internal/commands/delete"
	"github.com/rkuprov/mbot/cmd/cli/internal/commands/update"
	"github.com/rkuprov/mbot/cmd/cli/internal/commands/view"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Options struct {
	GetToken auth.Cmd   `cmd:"" help:"authenticate a user and grant them a session token"`
	Add      add.Cmd    `cmd:"" help:"Add various entities to the database"`
	View     view.Cmd   `cmd:"" help:"Examine various entities in the database"`
	Update   update.Cmd `cmd:"" help:"Update various entities in the database"`
	Delete   delete.Cmd `cmd:"" help:"Delete various entities in the database"`
}

func main() {
	ctx := context.Background()
	client := mbotpbconnect.NewMBotServerServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithInterceptors(WithTokenInterceptor()),
	)

	cli := kong.Parse(&Options{},
		kong.Name("mbot"),
		kong.Description("A CLI for managing the mbot service"),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.BindTo(client, (*mbotpbconnect.MBotServerServiceClient)(nil)),
	)
	err := cli.Run(ctx)
	cli.FatalIfErrorf(err)
}

func WithTokenInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("token", os.Getenv("MBOT_SESSION_TOKEN"))
			req.Header().Set("id", os.Getenv("MBOT_USER_ID"))

			resp, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			err = os.Setenv("MBOT_SESSION_TOKEN", resp.Header().Get("token"))
			if err != nil {
				return nil, err
			}
			err = os.Setenv("MBOT_USER_ID", resp.Header().Get("id"))
			if err != nil {
				return nil, err
			}

			return resp, nil
		}
	}
}
