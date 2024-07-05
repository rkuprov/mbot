package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrTokenNotFound   = errors.New("token not found")
	ErrTokenExpired    = errors.New("token expired")
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

func (a *Auth) ConfirmAndRotateToken(ctx context.Context, token SessionToken) (SessionToken, error) {
	tx, err := a.pg.Begin(ctx)
	if err != nil {
		return SessionToken{}, err
	}
	defer tx.Rollback(ctx)
	err = tx.QueryRow(ctx, `
	SELECT
		expires_at
	FROM session 
	WHERE token = $1 and is_valid = true 
	`, token.Token).Scan(&token.ValidUntil)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SessionToken{}, ErrTokenNotFound
		}
		return SessionToken{}, err
	}
	if time.Now().UTC().After(token.ValidUntil) {
		fmt.Println("token expired")
		fmt.Println(time.Now().Local().String())
		fmt.Println(token.ValidUntil.Local().String())
		return SessionToken{}, ErrTokenExpired
	}
	_, err = tx.Exec(ctx, `
	UPDATE session
	SET is_valid = false
	WHERE token = $1
	`, token.Token)
	if err != nil {
		return SessionToken{}, err
	}
	newToken := newSessionToken(token.UserID)
	_, err = tx.Exec(ctx, `
	INSERT INTO session (
		user_id,
		token,
		is_valid,
		expires_at
	) VALUES ($1, $2, $3, $4)
	`,
		newToken.UserID,
		newToken.Token,
		newToken.IsValid,
		newToken.ValidUntil,
	)
	if err != nil {
		return SessionToken{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return SessionToken{}, err
	}

	return newToken, nil
}

func newSessionToken(id string) SessionToken {
	return SessionToken{
		UserID:     id,
		Token:      uuid.New().String(),
		IsValid:    true,
		ValidUntil: time.Now().UTC().Add(30 * time.Minute),
	}
}
