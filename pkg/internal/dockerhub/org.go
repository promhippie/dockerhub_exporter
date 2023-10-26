package dockerhub

import (
	"context"
	"fmt"
)

// ByOrg fetches repositories by organization name.
func (c *Client) ByOrg(ctx context.Context, name string) ([]*Repository, error) {
	if err := c.Refresh(ctx); err != nil {
		return nil, fmt.Errorf("failed to refresh auth: %w", err)
	}

	path := fmt.Sprintf("/v2/repositories/%s/", name)
	records := make([]*Repository, 0)

	for {
		req, err := c.NewRequest(ctx, "GET", path, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to parse org request: %w", err)
		}

		result := &repositoryResponse{}

		if _, err := c.Do(req, result); err != nil {
			return nil, fmt.Errorf("failed to fetch org repos: %w", err)
		}

		records = append(
			records,
			result.Repositories...,
		)

		if result.Next == "" {
			break
		}

		path = result.Next
	}

	return records, nil
}
