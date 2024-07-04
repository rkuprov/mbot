package auth

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_ConfirmAndRotateToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	auth, done, err := NewTestAuth()
	require.NoError(t, err)
	defer done()

	fmt.Println("creating user")
	err = auth.NewUser(ctx, "testuser", "test")
	require.NoError(t, err)

	fmt.Println("logging in")
	token, err := auth.Login(ctx, "testuser", "test")
	require.NoError(t, err)
	token2, err := auth.ConfirmAndRotateToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEqual(t, token.Token, token2.Token)
}
