package main

import (
	"context"
	"net/http"

	"github.com/alecthomas/kong"

	"github.com/rkuprov/mbot/cmd/cli/internal/commands/auth"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Options struct {
	GetToken auth.Cmd `cmd:"" help:"Authenticate a user and grant them a session token"`
}

func main() {
	ctx := context.Background()
	client := mbotpbconnect.NewMBotServerServiceClient(http.DefaultClient, "http://localhost:8080")

	cli := kong.Parse(&Options{},
		kong.Name("mbot"),
		kong.Description("A CLI for managing the mbot service"),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.BindTo(client, (*mbotpbconnect.MBotServerServiceClient)(nil)),
	)
	err := cli.Run(ctx)
	cli.FatalIfErrorf(err)
}
