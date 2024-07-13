package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_ConfirmAndRotateToken(t *testing.T) {
	ctx := context.Background()
	auth, _, err := NewTestAuth()
	require.NoError(t, err)

	err = auth.NewUser(ctx, "testuser", "test")
	require.NoError(t, err)

	token, err := auth.Login(ctx, "testuser", "test")
	require.NoError(t, err)
	token2, err := auth.ConfirmAndRotateToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEqual(t, token, token2)
}
