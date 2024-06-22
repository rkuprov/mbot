package add

type Cmd struct {
	Customer     Customer     `cmd:"" help:"Add a customer"`
	Subscription Subscription `cmd:"" help:"Add a subscription"`
}
