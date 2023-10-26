package dockerhub

import (
	"time"
)

// Repository defines the repository structure from the DockerHub API.
type Repository struct {
	Namespace     string    `json:"namespace"`
	Name          string    `json:"name"`
	Status        int       `json:"status"`
	Stars         int       `json:"star_count"`
	Pulls         int       `json:"pull_count"`
	Collaborators int       `json:"collaborator_count"`
	Starred       bool      `json:"has_starred"`
	Private       bool      `json:"is_private"`
	Automated     bool      `json:"is_automated"`
	Editable      bool      `json:"can_edit"`
	Migrated      bool      `json:"is_migrated"`
	Updated       time.Time `json:"last_updated"`
}

// repositoryResponse defines the structure for repository listings within DockerHub API.
type repositoryResponse struct {
	Count        int           `json:"count"`
	Next         string        `json:"next"`
	Prev         string        `json:"previous"`
	Repositories []*Repository `json:"results"`
}
