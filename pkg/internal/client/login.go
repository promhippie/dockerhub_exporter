package client

import (
	"bytes"
	"context"
	"encoding/json"
)

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
	if c.Token == "" {
		return c.Login(ctx)
	}

	// TODO: validate or renew token like https://pkg.go.dev/github.com/dgrijalva/jwt-go#Parser.ParseUnverified and https://stackoverflow.com/questions/58441574/how-to-parse-the-expiration-date-of-a-jwt-to-a-time-time-in-go

	return nil
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
