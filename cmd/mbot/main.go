package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/alecthomas/kong"

	"github.com/rkuprov/mbot/cmd/mbot/internal/commands/add"
	"github.com/rkuprov/mbot/cmd/mbot/internal/commands/auth"
	"github.com/rkuprov/mbot/cmd/mbot/internal/commands/delete"
	"github.com/rkuprov/mbot/cmd/mbot/internal/commands/update"
	"github.com/rkuprov/mbot/cmd/mbot/internal/commands/view"
	"github.com/rkuprov/mbot/cmd/mbot/internal/middleware"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Options struct {
	Login    auth.LoginCmd    `cmd:"" help:"authenticate a user and grant them a session token"`
	Logout   auth.LogoutCmd   `cmd:"" help:"revoke a user's session token"`
	Register auth.RegisterCmd `cmd:"" help:"register a new user"`

	Add    add.Cmd    `cmd:"" help:"Add various entities to the database"`
	View   view.Cmd   `cmd:"" help:"Examine various entities in the database"`
	Update update.Cmd `cmd:"" help:"Update various entities in the database"`
	Delete delete.Cmd `cmd:"" help:"Delete various entities in the database"`
}

func main() {
	ctx := context.Background()
	client := mbotpbconnect.NewMBotServerServiceClient(
		http.DefaultClient,
		"http://136.36.82.165:7777",
		connect.WithInterceptors(middleware.WithTokenInterceptor()),
		connect.WithGRPC(),
	)
	authClient := mbotpbconnect.NewMbotAuthServerServiceClient(
		http.DefaultClient,
		"http://136.36.82.165:7777",
		connect.WithGRPC(),
	)

	cli := kong.Parse(&Options{},
		kong.Name("mbot"),
		kong.Description("A CLI for managing the mbot service"),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.BindTo(client, (*mbotpbconnect.MBotServerServiceClient)(nil)),
		kong.BindTo(authClient, (*mbotpbconnect.MbotAuthServerServiceClient)(nil)),
	)
	err := cli.Run(ctx)
	cli.FatalIfErrorf(err)
}
