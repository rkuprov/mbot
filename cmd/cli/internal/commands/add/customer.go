package add

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Customer struct {
	Name    string `help:"Customer name"`
	Contact string `help:"Customer contact"`
	Email   string `help:"Customer email"`
}

func (c *Customer) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.CreateCustomer(ctx, &connect.Request[mbotpb.CreateCustomerRequest]{
		Msg: &mbotpb.CreateCustomerRequest{
			Name:    c.Name,
			Email:   c.Email,
			Contact: c.Contact,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(resp.Msg.GetMessage())
	return nil
}
