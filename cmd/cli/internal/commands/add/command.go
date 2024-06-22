package add

type Cmd struct {
	Customer Customer `cmd:"" help:"Add a customer"`
}
