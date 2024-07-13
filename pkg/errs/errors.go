package errs

import (
	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	SessionTokenKey = "session_token"
)

func WithSessionTokenDetail(token string) *connect.ErrorDetail {
	d := &errdetails.ErrorInfo{
		Metadata: map[string]string{
			SessionTokenKey: token,
		},
	}
	out, err := connect.NewErrorDetail(d)
	if err != nil {
		panic(err)
	}
	return out
}

func HandleServerError(code connect.Code, err error, details ...*connect.ErrorDetail) error {
	ret := connect.NewError(code, err)
	for _, d := range details {
		ret.AddDetail(d)
	}
	return ret
}
