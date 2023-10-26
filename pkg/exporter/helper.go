package exporter

import (
	"github.com/promhippie/dockerhub_exporter/pkg/internal/dockerhub"
)

func boolToFloat64(val bool) float64 {
	if val {
		return 1.0
	}

	return 0.0
}

func appendRepo(slice []*dockerhub.Repository, i *dockerhub.Repository) []*dockerhub.Repository {
	if i == nil {
		return slice
	}

	for _, el := range slice {
		if el.Namespace == i.Namespace && el.Name == i.Name {
			return slice
		}
	}

	return append(slice, i)
}
