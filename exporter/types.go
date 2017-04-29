package exporter

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"
)

const (
	// timeFormat is required to parse the API timestamps properly.
	timeFormat = "2006-01-02T15:04:05Z"
)

// Repo represents the API repo records.
type Repo struct {
	Owner     string     `json:"namespace"`
	Name      string     `json:"name"`
	Pulls     float64    `json:"pull_count"`
	Stars     float64    `json:"star_count"`
	Status    float64    `json:"status"`
	Automated bool       `json:"is_automated"`
	UpdatedAt CustomTime `json:"last_updated"`
}

// Key generates a usable map key for the repo.
func (r *Repo) Key() string {
	return path.Join(r.Owner, r.Name)
}

// Fetch gathers the repo content from the API.
func (r *Repo) Fetch(name string) error {
	res, err := simpleClient().Get(
		fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/", name),
	)

	if err != nil {
		return fmt.Errorf("failed to request %s repository. %s", name, err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return fmt.Errorf("failed to parse %s repository. %s", name, err)
	}

	return nil
}

// Collection represents the API response for lists.
type Collection struct {
	Next  string
	Repos []*Repo `json:"results"`
}

// Fetch gathers the collection content from API.
func (c *Collection) Fetch(name string) error {
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/", name)

	if err := c.pagination(name, url); err != nil {
		return err
	}

	return nil
}

// pagination fetches the records per page from the API.
func (c *Collection) pagination(name, url string) error {
	var (
		newCollection = &Collection{}
	)

	res, err := simpleClient().Get(
		url,
	)

	if err != nil {
		return fmt.Errorf("failed to request %s organization. %s", name, err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(newCollection); err != nil {
		return fmt.Errorf("failed to parse %s organization. %s", name, err)
	}

	c.Repos = append(
		c.Repos,
		newCollection.Repos...,
	)

	if newCollection.Next != "" {
		if err := c.pagination(name, newCollection.Next); err != nil {
			return err
		}
	}

	return nil
}

// CustomTime represents the custom time format from the API.
type CustomTime struct {
	time.Time
}

// UnmarshalJSON properly unmarshals the time from JSON.
func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(timeFormat, strings.Replace(string(b), "\"", "", -1))

	if err != nil {
		t.Time = time.Time{}
	}

	return err
}
