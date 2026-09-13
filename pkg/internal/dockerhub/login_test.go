package dockerhub

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func token(t *testing.T, claims TokenClaims) string {
	t.Helper()

	payload, err := json.Marshal(claims)
	assert.NoError(t, err)

	return fmt.Sprintf(
		"header.%s.signature",
		base64.RawURLEncoding.EncodeToString(payload),
	)
}

func TestTokenExpiredWithValidToken(t *testing.T) {
	expired, err := tokenExpired(
		token(t, TokenClaims{
			Exp: time.Now().Add(10 * time.Minute).Unix(),
		}),
	)

	assert.NoError(t, err)
	assert.False(t, expired)
}

func TestTokenExpiredWithExpiredToken(t *testing.T) {
	expired, err := tokenExpired(
		token(t, TokenClaims{
			Exp: time.Now().Add(-1 * time.Minute).Unix(),
		}),
	)

	assert.NoError(t, err)
	assert.True(t, expired)
}

func TestTokenExpiredWithinLeeway(t *testing.T) {
	expired, err := tokenExpired(
		token(t, TokenClaims{
			Exp: time.Now().Add(expiryLeeway / 2).Unix(),
		}),
	)

	assert.NoError(t, err)
	assert.True(t, expired)
}

func TestTokenExpiredWithoutExpiry(t *testing.T) {
	expired, err := tokenExpired(
		token(t, TokenClaims{}),
	)

	assert.Error(t, err)
	assert.True(t, expired)
}

func TestTokenExpiredWithMalformedToken(t *testing.T) {
	for _, value := range []string{
		"",
		"not-a-token",
		"header.signature",
		"header.!!!not-base64!!!.signature",
		"header.bm90LWpzb24.signature",
	} {
		expired, err := tokenExpired(value)

		assert.Error(t, err, value)
		assert.True(t, expired, value)
	}
}

func TestRefreshWithoutCredentials(t *testing.T) {
	client := &Client{}

	assert.NoError(t, client.Refresh(t.Context()))
	assert.Empty(t, client.Token)
}
