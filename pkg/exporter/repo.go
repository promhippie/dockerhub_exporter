package exporter

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/promhippie/dockerhub_exporter/pkg/config"
	"github.com/promhippie/dockerhub_exporter/pkg/internal/client"
)

// RepoCollector collects metrics about the servers.
type RepoCollector struct {
	client   *client.Client
	logger   log.Logger
	failures *prometheus.CounterVec
	duration *prometheus.HistogramVec
	config   config.Target

	Status        *prometheus.Desc
	Stars         *prometheus.Desc
	Pulls         *prometheus.Desc
	Collaborators *prometheus.Desc
	Starred       *prometheus.Desc
	Private       *prometheus.Desc
	Automated     *prometheus.Desc
	Editable      *prometheus.Desc
	Migrated      *prometheus.Desc
	Updated       *prometheus.Desc
}

// NewRepoCollector returns a new RepoCollector.
func NewRepoCollector(logger log.Logger, c *client.Client, failures *prometheus.CounterVec, duration *prometheus.HistogramVec, cfg config.Target) *RepoCollector {
	if failures != nil {
		failures.WithLabelValues("repo").Add(0)
	}

	labels := []string{"owner", "name"}
	return &RepoCollector{
		client:   c,
		logger:   log.With(logger, "collector", "repo"),
		failures: failures,
		duration: duration,
		config:   cfg,

		Status: prometheus.NewDesc(
			"dockerhub_repo_status",
			"What is the current status of the repository",
			labels,
			nil,
		),
		Stars: prometheus.NewDesc(
			"dockerhub_repo_stars",
			"How often have this repository been stared",
			labels,
			nil,
		),
		Pulls: prometheus.NewDesc(
			"dockerhub_repo_pulls",
			"How often have this repository been pulled",
			labels,
			nil,
		),
		Collaborators: prometheus.NewDesc(
			"dockerhub_repo_collaborators",
			"How many collaborators are working on that repo",
			labels,
			nil,
		),
		Private: prometheus.NewDesc(
			"dockerhub_repo_private",
			"Is the repository private or public",
			labels,
			nil,
		),
		Automated: prometheus.NewDesc(
			"dockerhub_repo_automated",
			"Defines if the repository builds automatically",
			labels,
			nil,
		),
		Migrated: prometheus.NewDesc(
			"dockerhub_repo_migrated",
			"Hsve this repository been migrated",
			labels,
			nil,
		),
		Updated: prometheus.NewDesc(
			"dockerhub_repo_updated",
			"A timestamp when the repository have been updated",
			labels,
			nil,
		),
	}
}

// Metrics simply returns the list metric descriptors for generating a documentation.
func (c *RepoCollector) Metrics() []*prometheus.Desc {
	return []*prometheus.Desc{
		c.Status,
		c.Stars,
		c.Pulls,
		c.Collaborators,
		c.Private,
		c.Automated,
		c.Migrated,
		c.Updated,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *RepoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Status
	ch <- c.Stars
	ch <- c.Pulls
	ch <- c.Collaborators
	ch <- c.Private
	ch <- c.Automated
	ch <- c.Migrated
	ch <- c.Updated
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *RepoCollector) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()

	now := time.Now()
	repos := make([]*client.Repository, 0)

	for _, org := range c.config.Orgs.Value() {
		result, err := c.client.ByOrg(ctx, org)

		if err != nil {
			level.Error(c.logger).Log(
				"msg", "Failed to fetch repos",
				"type", "org",
				"name", org,
				"err", err,
			)

			c.failures.WithLabelValues("repo").Inc()
			continue
		}

		for _, repo := range result {
			repos = appendRepo(
				repos,
				repo,
			)
		}
	}

	for _, user := range c.config.Users.Value() {
		result, err := c.client.ByUser(ctx, user)

		if err != nil {
			level.Error(c.logger).Log(
				"msg", "Failed to fetch repos",
				"type", "user",
				"name", user,
				"err", err,
			)

			c.failures.WithLabelValues("repo").Inc()
			continue
		}

		for _, repo := range result {
			repos = appendRepo(
				repos,
				repo,
			)
		}
	}

	for _, repo := range c.config.Repos.Value() {
		result, err := c.client.ByName(ctx, repo)

		if err != nil {
			level.Error(c.logger).Log(
				"msg", "Failed to fetch repos",
				"type", "repo",
				"name", repo,
				"err", err,
			)

			c.failures.WithLabelValues("repo").Inc()
			continue
		}

		for _, repo := range result {
			repos = appendRepo(
				repos,
				repo,
			)
		}
	}

	c.duration.WithLabelValues("repo").Observe(time.Since(now).Seconds())

	level.Debug(c.logger).Log(
		"msg", "Fetched repos",
		"count", len(repos),
		"duration", time.Since(now),
	)

	for _, repo := range repos {
		labels := []string{
			repo.Namespace,
			repo.Name,
		}

		ch <- prometheus.MustNewConstMetric(
			c.Status,
			prometheus.GaugeValue,
			float64(repo.Status),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Stars,
			prometheus.GaugeValue,
			float64(repo.Stars),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Pulls,
			prometheus.GaugeValue,
			float64(repo.Pulls),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Collaborators,
			prometheus.GaugeValue,
			float64(repo.Collaborators),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Private,
			prometheus.GaugeValue,
			boolToFloat64(repo.Private),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Automated,
			prometheus.GaugeValue,
			boolToFloat64(repo.Automated),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Migrated,
			prometheus.GaugeValue,
			boolToFloat64(repo.Migrated),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Updated,
			prometheus.GaugeValue,
			float64(repo.Updated.Unix()),
			labels...,
		)
	}
}
