package auth

import (
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/mbot/internal/ui"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type LoginCmd struct {
	User string `arg:"" group:"Auth:" help:"Username" required:""`
	Pass string `arg:"" group:"Auth:" help:"Password" required:""`
}

func (c *LoginCmd) Run(ctx context.Context, client mbotpbconnect.MbotAuthServerServiceClient) error {
	resp, err := client.Login(ctx, &connect.Request[mbotpb.LoginRequest]{
		Msg: &mbotpb.LoginRequest{
			Username: c.User,
			Password: c.Pass,
		},
	})
	if err != nil {
		return err
	}
	f, err := os.Create(auth.SessionFile)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%s", resp.Header().Get(auth.HeaderSessionToken))

	ui.Tabular(ui.PrintCfg{
		Title: "Login Successful",
		Body:  []table.Row{{fmt.Sprintf("User %s was authenticated", c.User)}},
	})
	return nil
}
