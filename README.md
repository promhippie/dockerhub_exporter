# Docker Hub Exporter

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/dockerhub_exporter/status.svg)](http://github.dronehippie.de/webhippie/dockerhub_exporter)
[![Go Doc](https://godoc.org/github.com/webhippie/dockerhub_exporter?status.svg)](http://godoc.org/github.com/webhippie/dockerhub_exporter)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/dockerhub_exporter)](http://goreportcard.com/report/github.com/webhippie/dockerhub_exporter)
[![](https://images.microbadger.com/badges/image/tboerger/dockerhub-exporter.svg)](http://microbadger.com/images/tboerger/dockerhub-exporter "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/webhippie/general](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/webhippie/general)

A [Prometheus](https://prometheus.io/) exporter that collects Docker Hub statistics for defined namespaces and repositories.


## Installation

If you are missing something just write us on our nice [Gitter](https://gitter.im/webhippie/general) chat. If you find a security issue please contact thomas@webhippie.de first. Currently we are providing only a Docker image at `tboerger/dockerhub-exporter`.


### Usage

```bash
# docker run -ti --rm tboerger/dockerhub-exporter -h
Usage of /bin/dockerhub_exporter:
  -dockerhub.org value
      Organizations to watch on Docker Hub
  -dockerhub.repo value
      Repositories to watch on Docker Hub
  -log.format value
      Set the log target and format. Example: "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true" (default "logger:stderr")
  -log.level value
      Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal] (default "info")
  -version
      Print version information
  -web.listen-address string
      Address to listen on for web interface and telemetry (default ":9505")
  -web.telemetry-path string
      Path to expose metrics of the exporter (default "/metrics")
```


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). It is also possible to just simply execute the `go get github.com/webhippie/dockerhub_exporter` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/dockerhub_exporter
cd $GOPATH/src/github.com/webhippie/dockerhub_exporter
make test build

./dockerhub_exporter -h
```


## Metrics

```
# HELP dockerhub_automated Defines if the repository builds automatically
# TYPE dockerhub_automated gauge
dockerhub_automated{owner="tboerger",repo="redirects"} 0
# HELP dockerhub_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which dockerhub_exporter was built.
# TYPE dockerhub_exporter_build_info gauge
dockerhub_exporter_build_info{branch="HEAD",goversion="go1.8.1",revision="d1d5c9884f3d447a29348cad700c28758a8c146c",version="0.2.0"} 1
# HELP dockerhub_pulls How often have this repository been pulled
# TYPE dockerhub_pulls gauge
dockerhub_pulls{owner="tboerger",repo="redirects"} 6084
# HELP dockerhub_stars How often have this repository been stared
# TYPE dockerhub_stars gauge
dockerhub_stars{owner="tboerger",repo="redirects"} 0
# HELP dockerhub_status What is the current status of the repository
# TYPE dockerhub_status gauge
dockerhub_status{owner="tboerger",repo="redirects"} 1
# HELP dockerhub_up Check if Docker Hub response can be processed
# TYPE dockerhub_up gauge
dockerhub_up 1
# HELP dockerhub_updated A timestamp when the repository have been updated
# TYPE dockerhub_updated gauge
dockerhub_updated{owner="tboerger",repo="redirects"} 1.493281295e+09
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2017 Thomas Boerger <http://www.webhippie.de>
```
