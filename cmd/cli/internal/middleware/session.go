package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"connectrpc.com/connect"
)

const (
	HeaderSessionToken = "mbot-session-token"
	HeaderUserID       = "mbot-user-id"
	SessionFile        = ".session"
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
				req.Header().Get(HeaderUserID),
				req.Header().Get(HeaderSessionToken),
			)
			if err != nil {
				return nil, err
			}

			return resp, nil
		}
	}
}

func insertSessionToken(req connect.AnyRequest) (connect.AnyRequest, error) {
	f, err := os.OpenFile(SessionFile, os.O_RDWR, 0644)
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

	req.Header().Set(HeaderUserID, string(secrets[0]))
	req.Header().Set(HeaderSessionToken, string(secrets[1]))

	return req, nil
}

func updateSessionToken(id, token string) error {
	f, err := os.OpenFile(SessionFile, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	_, err = fmt.Fprintf(f, "%s\n%s", id, token)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
