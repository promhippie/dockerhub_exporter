package dockerhub

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// expiryLeeway defines how long before the actual expiry an access token is
// already considered expired, so that a request is not sent with a token that
// expires while it is in flight.
const expiryLeeway = 60 * time.Second

// Login is handling the login for the DockerHub API.
func (c *Client) Login(ctx context.Context) error {
	payload, err := json.Marshal(LoginRequest{
		Username: c.Username,
		Password: c.Password,
	})

	if err != nil {
		return err
	}

	req, err := c.NewRequest(
		ctx,
		"POST",
		"/v2/users/login",
		bytes.NewBuffer(
			payload,
		),
	)

	if err != nil {
		return err
	}

	result := LoginResponse{}

	if _, err := c.Do(req, &result); err != nil {
		return err
	}

	c.Token = result.Token

	return nil
}

// Refresh is handling refreshing the access token for DockerHub API.
func (c *Client) Refresh(ctx context.Context) error {
	if c.Username == "" || c.Password == "" {
		// Without credentials there is nothing to refresh, the public
		// endpoints are served fine for anonymous requests.
		return nil
	}

	if c.Token == "" {
		return c.Login(ctx)
	}

	expired, err := tokenExpired(c.Token)

	if err != nil || expired {
		return c.Login(ctx)
	}

	return nil
}

// tokenExpired checks if the access token is expired or about to expire. The
// signature is deliberately not verified, we only read the expiry claim to
// decide whether a new token has to be requested.
func tokenExpired(token string) (bool, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return true, fmt.Errorf("token is not a json web token")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])

	if err != nil {
		return true, fmt.Errorf("failed to decode token payload: %w", err)
	}

	claims := TokenClaims{}

	if err := json.Unmarshal(payload, &claims); err != nil {
		return true, fmt.Errorf("failed to parse token payload: %w", err)
	}

	if claims.Exp == 0 {
		return true, fmt.Errorf("token does not carry an expiry")
	}

	return time.Now().Add(expiryLeeway).After(
		time.Unix(claims.Exp, 0),
	), nil
}

// LoginRequest defines the request structure for login handler.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse defines the response structure for login handler.
type LoginResponse struct {
	Token string `json:"token"`
}

// TokenClaims defines the subset of the access token claims we care about.
type TokenClaims struct {
	Exp int64 `json:"exp"`
}
