package update

type Cmd struct {
	Customer Customer `cmd:"" help:"Update a customer"`
}
