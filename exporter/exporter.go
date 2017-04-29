package exporter

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const (
	// namespace defines the Prometheus namespace for this exporter.
	namespace = "dockerhub"
)

var (
	// isUp defines if the API response can get processed.
	isUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Check if Docker Hub response can be processed",
		},
	)

	// isAutomated defines a map to collect if the repository is an automated build.
	isAutomated = map[string]prometheus.Gauge{}

	// pullCount defines a map to collect the number of pulls per repository.
	pullCount = map[string]prometheus.Gauge{}

	// starCount defines a map to collect the number of stars per repository.
	starCount = map[string]prometheus.Gauge{}

	// currentStatus defines a map to collect the status of the repositories.
	currentStatus = map[string]prometheus.Gauge{}

	// updatedAt defines a map to collect the last update timestamp per repository.
	updatedAt = map[string]prometheus.Gauge{}
)

// init just defines the initial state of the exports.
func init() {
	isUp.Set(0)
}

// NewExporter gives you a new exporter instance.
func NewExporter(orgs, repos []string) *Exporter {
	return &Exporter{
		orgs:  orgs,
		repos: repos,
	}
}

// Exporter combines the metric collector and descritions.
type Exporter struct {
	orgs  []string
	repos []string
	mutex sync.RWMutex
}

// Describe defines the metric descriptions for Prometheus.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- isUp.Desc()

	for _, metric := range isAutomated {
		ch <- metric.Desc()
	}

	for _, metric := range pullCount {
		ch <- metric.Desc()
	}

	for _, metric := range starCount {
		ch <- metric.Desc()
	}

	for _, metric := range currentStatus {
		ch <- metric.Desc()
	}

	for _, metric := range updatedAt {
		ch <- metric.Desc()
	}
}

// Collect delivers the metrics to Prometheus.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if err := e.scrape(); err != nil {
		log.Error(err)

		isUp.Set(0)
		ch <- isUp

		return
	}

	ch <- isUp

	for _, metric := range isAutomated {
		ch <- metric
	}

	for _, metric := range pullCount {
		ch <- metric
	}

	for _, metric := range starCount {
		ch <- metric
	}

	for _, metric := range currentStatus {
		ch <- metric
	}

	for _, metric := range updatedAt {
		ch <- metric
	}
}

// scrape just starts the scraping loop.
func (e *Exporter) scrape() error {
	log.Debug("start scrape loop")

	for _, org := range e.orgs {
		log.Debugf("checking %s organization", org)

		if err := processOrg(org); err != nil {
			return err
		}
	}

	for _, repo := range e.repos {
		log.Debugf("checking %s repository", repo)

		if err := processRepo(repo); err != nil {
			return err
		}
	}

	isUp.Set(1)
	return nil
}

// processOrg fetches the organization content from the API.
func processOrg(name string) error {
	var (
		collection = &Collection{}
	)

	if err := collection.Fetch(name); err != nil {
		log.Debugf("%s", err)
		return fmt.Errorf("failed to fetch %s organization", name)
	}

	for _, repo := range collection.Repos {
		if err := scrapeRepo(repo); err != nil {
			return err
		}
	}

	return nil
}

// processRepo fetches the repository content from the API.
func processRepo(name string) error {
	var (
		repo = &Repo{}
	)

	if err := repo.Fetch(name); err != nil {
		log.Debugf("%s", err)
		return fmt.Errorf("failed to fetch %s repository", name)
	}

	return scrapeRepo(repo)
}

// scrapeRepo processes the content of a specific repository.
func scrapeRepo(repo *Repo) error {
	if _, ok := isAutomated[repo.Key()]; ok == false {
		isAutomated[repo.Key()] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "automated",
				Help:      "Defines if the repository builds automatically",
				ConstLabels: prometheus.Labels{
					"owner": repo.Owner,
					"repo":  repo.Name,
				},
			},
		)
	}

	if _, ok := pullCount[repo.Key()]; ok == false {
		pullCount[repo.Key()] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "pulls",
				Help:      "How often have this repository been pulled",
				ConstLabels: prometheus.Labels{
					"owner": repo.Owner,
					"repo":  repo.Name,
				},
			},
		)
	}

	if _, ok := starCount[repo.Key()]; ok == false {
		starCount[repo.Key()] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "stars",
				Help:      "How often have this repository been stared",
				ConstLabels: prometheus.Labels{
					"owner": repo.Owner,
					"repo":  repo.Name,
				},
			},
		)
	}

	if _, ok := currentStatus[repo.Key()]; ok == false {
		currentStatus[repo.Key()] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "status",
				Help:      "What is the current status of the repository",
				ConstLabels: prometheus.Labels{
					"owner": repo.Owner,
					"repo":  repo.Name,
				},
			},
		)
	}

	if _, ok := updatedAt[repo.Key()]; ok == false {
		updatedAt[repo.Key()] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "updated",
				Help:      "A timestamp when the repository have been updated",
				ConstLabels: prometheus.Labels{
					"owner": repo.Owner,
					"repo":  repo.Name,
				},
			},
		)
	}

	if repo.Automated {
		isAutomated[repo.Key()].Set(1.0)
	} else {
		isAutomated[repo.Key()].Set(0.0)
	}

	pullCount[repo.Key()].Set(repo.Pulls)
	starCount[repo.Key()].Set(repo.Stars)
	currentStatus[repo.Key()].Set(repo.Status)
	updatedAt[repo.Key()].Set(float64(repo.UpdatedAt.Unix()))

	return nil
}
