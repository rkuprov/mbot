package ui

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

func Single(msg interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	var title string
	var text []string
	var header table.Row
	switch resp := msg.(type) {
	case *mbotpb.CreateSubscriptionResponse:
		if resp.GetSubscription() == nil {
			title = "Failure!"
			text = append(text, resp.GetMessage())
		}
		title = fmt.Sprintf("Success! Subscription created for customer ID: %s", resp.GetCustomerId())
		header = table.Row{"ID", "Start Date", "Expiration Date"}
		text = []string{resp.GetSubscription().SubscriptionId,
			resp.GetSubscription().GetStartDate().AsTime().Format("2006-01-02"),
			resp.GetSubscription().GetExpirationDate().AsTime().Format("2006-01-02")}
	case *mbotpb.CreateCustomerResponse:
		fmt.Println("CreateCustomerResponse")
		if resp.GetId() == "" {
			title = "Failure!"
			text = append(text, resp.GetMessage())
		}
		title = fmt.Sprintf("Success!")
		text = []string{resp.GetMessage()}
	}

	t.SetTitle(title)
	if len(header) > 0 {
		t.AppendHeader(header)
	}
	row := table.Row{}
	for _, txt := range text {
		row = append(row, txt)
	}
	t.AppendRow(row)
	t.Render()
}
