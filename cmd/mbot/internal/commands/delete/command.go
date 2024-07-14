package delete

type Cmd struct {
	Customer     Customer     `cmd:"" help:"Delete a customer"`
	Subscription Subscription `cmd:"" help:"Delete a subscription"`
}
