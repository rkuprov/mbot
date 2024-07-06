package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

var (
	ErrTokenNotFound   = errors.New("token not found")
	ErrTokenExpired    = errors.New("token expired")
	ErrTokenInvalid    = errors.New("token invalid")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidUsername = errors.New("invalid username")
	ErrUserNotFound    = errors.New("user not found")
)

type SessionToken struct {
	UserID     string
	Token      string
	IsValid    bool
	ValidUntil time.Time
}

func (a *Auth) ConfirmAndRotateToken(ctx context.Context, tokenVal string) (*mbotpb.SessionToken, error) {
	tx, err := a.pg.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	var exp time.Time
	var isValid bool
	err = tx.QueryRow(ctx, `
	SELECT
	    is_valid,
		expires_at
	FROM session 
	WHERE token = $1
	`, tokenVal).Scan(&isValid, &exp)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTokenNotFound
		}
		return nil, err
	}
	if time.Now().UTC().After(exp) {
		return nil, ErrTokenExpired
	}
	if !isValid {
		return nil, ErrTokenInvalid

	}
	_, err = tx.Exec(ctx, `
	UPDATE session
	SET is_valid = false
	WHERE token = $1
	`, tokenVal)
	if err != nil {
		return nil, err
	}
	newToken := newSessionToken()
	_, err = tx.Exec(ctx, `
	INSERT INTO session (
		token
	) VALUES ($1)
	`,
		newToken.Value,
		newToken.Expiration.AsTime(),
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func newSessionToken() *mbotpb.SessionToken {
	return &mbotpb.SessionToken{
		Value:      uuid.New().String(),
		Expiration: timestamppb.New(time.Now().UTC().Add(time.Hour * 24)),
	}
}
