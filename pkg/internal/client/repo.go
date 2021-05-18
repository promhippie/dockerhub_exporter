package client

import (
	"context"
	"fmt"
)

// ByName fetches repositories by repository name.
func (c *Client) ByName(ctx context.Context, name string) ([]*Repository, error) {
	if err := c.Refresh(ctx); err != nil {
		return nil, fmt.Errorf("failed to refresh auth: %w", err)
	}

	path := fmt.Sprintf("/v2/repositories/%s/", name)
	records := make([]*Repository, 0)

	req, err := c.NewRequest(ctx, "GET", path, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	result := &Repository{}

	if _, err := c.Do(req, result); err != nil {
		return nil, fmt.Errorf("failed to fetch repos: %w", err)
	}

	return append(
		records,
		result,
	), nil
}
