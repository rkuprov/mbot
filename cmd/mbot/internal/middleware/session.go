package middleware

import (
	"context"
	"io"
	"os"

	"connectrpc.com/connect"

	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/errs"
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
				return nil, errs.HandleClientError(err)
			}

			err = auth.UpdateSessionToken(
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
