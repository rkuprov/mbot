package auth

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/mbot/internal/ui"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type RegisterCmd struct {
	Username string `arg:"" group:"Auth:" help:"Username" required:""`
	Password string `arg:"" group:"Auth:" help:"Password" required:""`
}

func (c *RegisterCmd) Run(ctx context.Context, client mbotpbconnect.MbotAuthServerServiceClient) error {
	fmt.Printf("Please, re-enter your password to confirm:\n\n")
	pw, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return err
	}
	if c.Password != string(pw) {
		ui.Tabular(ui.PrintCfg{
			Body: []table.Row{{"Passwords do not match"}},
		})
		return nil
	}

	resp, err := client.Register(ctx, &connect.Request[mbotpb.RegisterRequest]{
		Msg: &mbotpb.RegisterRequest{
			Username: c.Username,
			Password: c.Password,
		},
	})
	if err != nil {
		return err
	}

	if !resp.Msg.Ok {
		ui.Tabular(ui.PrintCfg{
			Body: []table.Row{{"Failed to register user"}},
		})
	}

	fmt.Println()
	ui.Tabular(ui.PrintCfg{
		Title: "Registration Successful",
		Body:  []table.Row{{fmt.Sprintf("User %s was registered", c.Username)}},
	})

	return nil
}
