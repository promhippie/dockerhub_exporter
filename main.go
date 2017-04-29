package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/webhippie/dockerhub_exporter/exporter"

	_ "net/http/pprof"
)

var (
	// showVersion is a flag to display the current version.
	showVersion = flag.Bool("version", false, "Print version information")

	// listenAddress defines the local address binding for the server.
	listenAddress = flag.String("web.listen-address", ":9105", "Address to listen on for web interface and telemetry")

	// metricsPath defines the path to access the metrics.
	metricsPath = flag.String("web.telemetry-path", "/metrics", "Path to expose metrics of the exporter")

	// orgs defines the organizations to export.
	orgs StringSlice

	// repos defines the repositories to export.
	repos StringSlice
)

// init registers the collector version.
func init() {
	prometheus.MustRegister(version.NewCollector("dockerhub_exporter"))
}

// main simply initializes this tool.
func main() {
	flag.Var(&orgs, "dockerhub.org", "Organizations to watch on Docker Hub")
	flag.Var(&repos, "dockerhub.repo", "Repositories to watch on Docker Hub")

	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("dockerhub_exporter"))
		os.Exit(0)
	}

	log.Infoln("Starting Docker Hub exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	e := exporter.NewExporter(orgs, repos)

	prometheus.MustRegister(e)
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(os.Getpid(), ""))

	http.Handle(*metricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Infof("Listening on %s", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}

// StringSlice represents a custom string slice flag.
type StringSlice []string

// String represents the string slice as a string.
func (ss *StringSlice) String() string {
	return strings.Join(*ss, ",")
}

// Set appends the string slice value to the current list.
func (ss *StringSlice) Set(value string) error {
	if value != "" {
		*ss = append(*ss, strings.Split(strings.Replace(value, " ", "", -1), ",")...)
	}

	return nil
}
