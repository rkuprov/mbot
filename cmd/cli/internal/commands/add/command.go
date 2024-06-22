package add

type Cmd struct {
	Customer     Customer     `cmd:"" help:"Add a customer"`
	Subscription Subscription `cmd:"" aliases:"sub" help:"Add a subscription"`
}
