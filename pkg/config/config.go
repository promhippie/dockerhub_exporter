package config

import (
	"time"

	"github.com/urfave/cli/v2"
)

// Server defines the general server configuration.
type Server struct {
	Addr    string
	Path    string
	Timeout time.Duration
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
}

// Target defines the target specific configuration.
type Target struct {
	Username string
	Password string
	Orgs     cli.StringSlice
	Users    cli.StringSlice
	Repos    cli.StringSlice
	Timeout  time.Duration
}

// Collector defines the collector specific configuration.
type Collector struct {
	Orgs  bool
	Repos bool
}

// Config is a combination of all available configurations.
type Config struct {
	Server    Server
	Logs      Logs
	Target    Target
	Collector Collector
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
