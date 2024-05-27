package datamodel

import "fmt"

type Customer struct {
	ID      string
	Name    string
	Email   string
	Contact string
	Token   string
}

func (c Customer) Encode() []byte {
	in := fmt.Sprintf("%d#%s#%s#%s#%s", c.ID, c.Name, c.Email, c.Contact, c.Token)
	return []byte(in)
}
