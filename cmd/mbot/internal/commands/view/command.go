package view

type Cmd struct {
	Customer     Customer     `cmd:"" help:"View a customer"`
	Subscription Subscription `cmd:"" aliases:"sub" help:"View a subscription"`
}
