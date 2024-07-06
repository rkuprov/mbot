package auth

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost   = 10
	HeaderSessionToken = "mbot-session-token"
	HeaderUserID       = "mbot-user-id"
	SessionFile        = ".session"
)

func (a *Auth) Login(ctx context.Context, username, password string) (string, error) {
	err := a.authenticate(ctx, username, password)
	if err != nil {
		return "", err
	}
	token := newSessionToken()

	_, err = a.pg.Exec(ctx, `
		INSERT INTO session (
			token
		) VALUES ($1)
	`,
		token.Value,
	)
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

func (a *Auth) authenticate(ctx context.Context, inUsrName, password string) error {
	var username, pw string
	dbErr := a.pg.QueryRow(ctx, `
	SELECT
	username,
	password
	FROM users
	WHERE username = $1
	`, inUsrName).Scan(&username, &pw)
	if dbErr != nil && !errors.Is(dbErr, pgx.ErrNoRows) {
		return dbErr
	}
	dbPwHash, err := stringToHash(pw)
	if err != nil {
		return err
	}

	switch {
	case username != inUsrName:
		return ErrInvalidUsername
	case bcrypt.CompareHashAndPassword(dbPwHash, []byte(password)) != nil:
		return ErrInvalidPassword
	case errors.Is(dbErr, pgx.ErrNoRows):
		return ErrUserNotFound
	}

	return nil
}

func (a *Auth) NewUser(ctx context.Context, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	if err != nil {
		return err
	}
	_, err = a.pg.Exec(ctx, `
		INSERT INTO users (
			username,
			password
		) VALUES ($1, $2)
	`, username, hashToString(hash))

	if err != nil {
		return err
	}
	return nil
}

func stringToHash(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
func hashToString(h []byte) string {
	return base64.URLEncoding.EncodeToString(h)
}

func (a *Auth) Logout(ctx context.Context, token SessionToken) error {
	var present bool
	err := a.pg.QueryRow(ctx, `
	SELECT
	EXISTS (
		SELECT 1
		FROM session
		WHERE token = $1)
	`, token.Token).Scan(&present)
	if err != nil {
		return err
	}
	if !present {
		return ErrTokenNotFound
	}

	_, err = a.pg.Exec(ctx, `
	TRUNCATE TABLE session
	`)
	if err != nil {
		return err
	}

	return nil
}
