package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost = 10
)

func (a *Auth) Login(ctx context.Context, username, password string) (SessionToken, error) {
	id, err := a.authenticate(ctx, username, password)
	if err != nil {
		return SessionToken{}, err
	}
	fmt.Println("authenticated")
	token := newSessionToken(id)

	_, err = a.pg.Exec(ctx, `
		INSERT INTO session (
			user_id,
			token,
			is_valid,
			expires_at
		) VALUES ($1, $2, $3, $4)
	`,
		token.UserID,
		token.Token,
		token.IsValid,
		token.ValidUntil)
	if err != nil {
		return SessionToken{}, err
	}

	return token, nil
}

func (a *Auth) authenticate(ctx context.Context, inUsrName, password string) (string, error) {
	var id, username, pw string
	dbErr := a.pg.QueryRow(ctx, `
	SELECT
	id,
	username,
	password
	FROM users
	WHERE username = $1
	`, inUsrName).Scan(&id, &username, &pw)
	if dbErr != nil && !errors.Is(dbErr, pgx.ErrNoRows) {
		return "", dbErr
	}
	dbPwHash, err := stringToHash(pw)
	if err != nil {
		return "", err
	}

	switch {
	case username != inUsrName:
		return "", ErrInvalidUsername
	case bcrypt.CompareHashAndPassword(dbPwHash, []byte(password)) != nil:
		return "", ErrInvalidPassword
	case errors.Is(dbErr, pgx.ErrNoRows):
		return "", ErrUserNotFound
	}

	return id, nil
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
	fmt.Println("user created")
	return nil
}

func stringToHash(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
func hashToString(h []byte) string {
	return base64.URLEncoding.EncodeToString(h)
}