package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/promhippie/dockerhub_exporter/pkg/action"
	"github.com/promhippie/dockerhub_exporter/pkg/config"
	"github.com/promhippie/dockerhub_exporter/pkg/version"
	"github.com/urfave/cli/v3"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.Command{
		Name:    "dockerhub_exporter",
		Version: version.String,
		Usage:   "DockerHub Exporter",
		Authors: []any{
			"Thomas Boerger <thomas@webhippie.de>",
		},
		Flags: RootFlags(cfg),
		Commands: []*cli.Command{
			Health(cfg),
		},
		Action: func(_ context.Context, _ *cli.Command) error {
			logger := setupLogger(cfg)

			if len(cfg.Target.Orgs) == 0 &&
				len(cfg.Target.Users) == 0 &&
				len(cfg.Target.Repos) == 0 {
				logger.Error("Missing required org, user or repo")
				return fmt.Errorf("missing required org, user or repo")
			}

			return action.Server(cfg, logger)
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	return app.Run(context.Background(), os.Args)
}

// RootFlags defines the available root flags.
func RootFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "info",
			Usage:       "Only log messages with given severity",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_LOG_LEVEL"),
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log.pretty",
			Value:       false,
			Usage:       "Enable pretty messages for logging",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_LOG_PRETTY"),
			Destination: &cfg.Logs.Pretty,
		},
		&cli.StringFlag{
			Name:        "web.address",
			Value:       "0.0.0.0:9505",
			Usage:       "Address to bind the metrics server",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_WEB_ADDRESS"),
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "web.path",
			Value:       "/metrics",
			Usage:       "Path to bind the metrics server",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_WEB_PATH"),
			Destination: &cfg.Server.Path,
		},
		&cli.BoolFlag{
			Name:        "web.debug",
			Value:       false,
			Usage:       "Enable pprof debugging for server",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_WEB_PPROF"),
			Destination: &cfg.Server.Pprof,
		},
		&cli.DurationFlag{
			Name:        "web.timeout",
			Value:       10 * time.Second,
			Usage:       "Server metrics endpoint timeout",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_WEB_TIMEOUT"),
			Destination: &cfg.Server.Timeout,
		},
		&cli.StringFlag{
			Name:        "web.config",
			Value:       "",
			Usage:       "Path to web-config file",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_WEB_CONFIG"),
			Destination: &cfg.Server.Web,
		},
		&cli.DurationFlag{
			Name:        "request.timeout",
			Value:       5 * time.Second,
			Usage:       "Timeout requesting DockerHub API",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_REQUEST_TIMEOUT"),
			Destination: &cfg.Target.Timeout,
		},
		&cli.StringFlag{
			Name:        "dockerhub.username",
			Value:       "",
			Usage:       "Username for the DockerHub authentication",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_USERNAME"),
			Destination: &cfg.Target.Username,
		},
		&cli.StringFlag{
			Name:        "dockerhub.password",
			Value:       "",
			Usage:       "Password for the DockerHub authentication",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_PASSWORD"),
			Destination: &cfg.Target.Password,
		},
		&cli.StringSliceFlag{
			Name:        "dockerhub.org",
			Value:       []string{},
			Usage:       "Organizations to scrape metrics from",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_ORG"),
			Destination: &cfg.Target.Orgs,
		},
		&cli.StringSliceFlag{
			Name:        "dockerhub.user",
			Value:       []string{},
			Usage:       "Users to scrape metrics from",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_USER"),
			Destination: &cfg.Target.Users,
		},
		&cli.StringSliceFlag{
			Name:        "dockerhub.repo",
			Value:       []string{},
			Usage:       "Repositories to scrape metrics from",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_REPO"),
			Destination: &cfg.Target.Repos,
		},
		&cli.BoolFlag{
			Name:        "collector.repos",
			Value:       true,
			Usage:       "Enable collector for repos",
			Sources:     cli.EnvVars("DOCKERHUB_EXPORTER_COLLECTOR_REPOS"),
			Destination: &cfg.Collector.Repos,
		},
	}
}
