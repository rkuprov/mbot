package update

type Cmd struct {
	Customer     Customer     `cmd:"" aliases:"cust,c" help:"Update a customer"`
	Subscription Subscription `cmd:"" aliases:"Sub,s" help:"Update a subscription"`
}
