package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Cmd struct {
	User string `group:"Auth:" help:"Username" required:""`
	Pass string `group:"Auth:" help:"Password" required:""`
}

func (c *Cmd) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.GetCustomersAll(ctx, &connect.Request[mbotpb.GetCustomersAllRequest]{})
	if err != nil {
		return err
	}
	fmt.Println("got the response")
	for _, customer := range resp.Msg.GetCustomers() {
		fmt.Println(customer)
	}
	return nil
}
