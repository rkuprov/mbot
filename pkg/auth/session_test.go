package auth

import (
	"context"
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

	err = auth.NewUser(ctx, "testuser", "test")
	require.NoError(t, err)

	token, err := auth.Login(ctx, "testuser", "test")
	require.NoError(t, err)
	token2, err := auth.ConfirmAndRotateToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEqual(t, token.Token, token2.Token)
}
