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
	showVersion = flag.Bool("version", false, "Print version information")

	listenAddress = flag.String("web.listen-address", ":9105", "Address to listen on for web interface and telemetry")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path to expose metrics of the exporter")

	orgs  StringSlice
	repos StringSlice
)

func init() {
	prometheus.MustRegister(version.NewCollector("dockerhub_exporter"))
}

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

// StringSlice
type StringSlice []string

// String
func (ss *StringSlice) String() string {
	return strings.Join(*ss, ",")
}

// Set
func (ss *StringSlice) Set(value string) error {
	if value != "" {
		*ss = append(*ss, strings.Split(strings.Replace(value, " ", "", -1), ",")...)
	}

	return nil
}
