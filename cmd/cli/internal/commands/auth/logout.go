package auth

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"connectrpc.com/connect"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/cmd/cli/internal/ui"
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
	secrets := bytes.Split(bytes.TrimSpace(bts), []byte("\n"))
	if len(secrets) != 2 {
		return fmt.Errorf("expected 2 secrets, got %d", len(secrets))
	}
	req.Header().Set(auth.HeaderUserID, string(secrets[0]))
	req.Header().Set(auth.HeaderSessionToken, string(secrets[1]))

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
