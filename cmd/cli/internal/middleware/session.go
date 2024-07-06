package middleware

import (
	"context"
	"fmt"
	"io"
	"os"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/auth"
)

func WithTokenInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req, err := insertSessionToken(req)
			if err != nil {
				return nil, err
			}

			resp, err := next(ctx, req)
			if err != nil {
				return nil, err
			}

			err = updateSessionToken(
				resp.Header().Get(auth.HeaderSessionToken),
			)
			if err != nil {
				return nil, err
			}

			return resp, nil
		}
	}
}

func insertSessionToken(req connect.AnyRequest) (connect.AnyRequest, error) {
	f, err := os.OpenFile(auth.SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	req.Header().Set(auth.HeaderSessionToken, string(bts))

	return req, nil
}

func updateSessionToken(token string) error {
	f, err := os.OpenFile(auth.SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s", token)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
