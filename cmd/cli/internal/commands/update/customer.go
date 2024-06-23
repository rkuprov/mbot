package update

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
)

type Customer struct {
	ID      string `arg:"" required:"" help:"Customer ID"`
	Name    string `help:"Customer name"`
	Email   string `help:"Customer email"`
	Contact string `help:"Customer contact"`
}

func (c *Customer) Run(ctx context.Context, client mbotpbconnect.MBotServerServiceClient) error {
	resp, err := client.UpdateCustomer(ctx, &connect.Request[mbotpb.UpdateCustomerRequest]{
		Msg: &mbotpb.UpdateCustomerRequest{
			Id:      c.ID,
			Name:    c.Name,
			Email:   c.Email,
			Contact: c.Contact,
		},
	})
	if err != nil {
		return err
	}

	fmt.Println(resp.Msg)

	return nil
}
