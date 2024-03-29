package command

import (
	"testing"

	"github.com/promhippie/dockerhub_exporter/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	logger := setupLogger(config.Load())
	assert.NotNil(t, logger)
}
