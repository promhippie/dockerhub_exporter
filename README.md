# DockerHub Exporter

[![Current Tag](https://img.shields.io/github/v/tag/promhippie/dockerhub_exporter?sort=semver)](https://github.com/promhippie/prometheus-dockerhub-sd) [![General Build](https://github.com/promhippie/dockerhub_exporter/actions/workflows/general.yml/badge.svg)](https://github.com/promhippie/dockerhub_exporter/actions/workflows/general.yaml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/dc676e04911c448ebf06b03b5bf7ddaa)](https://www.codacy.com/gh/promhippie/dockerhub_exporter/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/dockerhub_exporter&amp;utm_campaign=Badge_Grade) [![Go Doc](https://godoc.org/github.com/promhippie/dockerhub_exporter?status.svg)](http://godoc.org/github.com/promhippie/dockerhub_exporter) [![Go Report](http://goreportcard.com/badge/github.com/promhippie/dockerhub_exporter)](http://goreportcard.com/report/github.com/promhippie/dockerhub_exporter)

An exporter for [Prometheus][prometheus] that collects metrics from
[DockerHub][docker].

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our containers published on [Docker Hub][dockerhub] and [Quay][quayio].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/promhippie/dockerhub_exporter.git
cd dockerhub_exporter

make generate build

./bin/dockerhub_exporter -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[prometheus]: https://prometheus.io
[docker]: https://hub.docker.com
[releases]: https://github.com/promhippie/dockerhub_exporter/releases
[dockerhub]: https://hub.docker.com/r/promhippie/dockerhub-exporter/tags/
[quayio]: https://quay.io/repository/promhippie/dockerhub-exporter?tab=tags
[docs]: https://promhippie.github.io/dockerhub_exporter/#getting-started
[golang]: http://golang.org/doc/install.html
