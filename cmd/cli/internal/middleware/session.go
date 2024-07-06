package middleware

import (
	"bytes"
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
			req, err := InsertSessionToken(req)
			if err != nil {
				return nil, err
			}

			resp, err := next(ctx, req)
			if err != nil {
				return nil, err
			}

			err = updateSessionToken(
				req.Header().Get(auth.HeaderUserID),
				req.Header().Get(auth.HeaderSessionToken),
			)
			if err != nil {
				return nil, err
			}

			return resp, nil
		}
	}
}

func InsertSessionToken(req connect.AnyRequest) (connect.AnyRequest, error) {
	f, err := os.OpenFile(auth.SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	secrets := bytes.Split(bytes.TrimSpace(bts), []byte("\n"))
	if len(secrets) != 2 {
		return nil, fmt.Errorf("expected 2 secrets, got %d", len(secrets))
	}

	req.Header().Set(auth.HeaderUserID, string(secrets[0]))
	req.Header().Set(auth.HeaderSessionToken, string(secrets[1]))

	return req, nil
}

func updateSessionToken(id, token string) error {
	f, err := os.OpenFile(auth.SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	_, err = fmt.Fprintf(f, "%s\n%s", id, token)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
