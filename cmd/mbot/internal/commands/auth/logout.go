package auth

import (
	"bytes"
	"context"
	"io"
	"os"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/mbot/internal/ui"
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type LogoutCmd struct{}

func (c *LogoutCmd) Run(ctx context.Context, client mbotpbconnect.MbotAuthServerServiceClient) error {
	req := connect.Request[mbotpb.LogoutRequest]{
		Msg: &mbotpb.LogoutRequest{},
	}

	f, err := os.OpenFile(auth.SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	req.Header().Set(auth.HeaderSessionToken, string(bytes.TrimSpace(bts)))

	_, err = client.Logout(ctx, &req)
	if err != nil {
		return err
	}
	err = f.Truncate(0)
	if err != nil {
		return err
	}
	ui.Tabular(ui.PrintCfg{
		Body: []table.Row{{"Logout Successful"}},
	})

	return nil
}
