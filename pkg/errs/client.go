package errs

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/rkuprov/mbot/pkg/auth"
)

func HandleClientError(err error) error {
	var connError *connect.Error
	if !errors.As(err, &connError) {
		return fmt.Errorf("could not extract errors: %w", err)
	}
	for _, detail := range connError.Details() {
		msg, valueErr := detail.Value()
		if valueErr != nil {
			return fmt.Errorf("could not extract value: %w", valueErr)
		}
		switch m := msg.(type) {
		case *errdetails.ErrorInfo:
			if token, ok := m.Metadata[SessionTokenKey]; ok {
				upateErr := auth.UpdateSessionToken(token)
				if upateErr != nil {
					return fmt.Errorf("could not update session token: %w: %w", upateErr, err)
				}
			}
			return err
		}
	}
	return err
}
