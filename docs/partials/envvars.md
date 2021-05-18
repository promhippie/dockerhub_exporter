DOCKERHUB_EXPORTER_LOG_LEVEL
: Only log messages with given severity, defaults to `info`

DOCKERHUB_EXPORTER_LOG_PRETTY
: Enable pretty messages for logging, defaults to `false`

DOCKERHUB_EXPORTER_WEB_ADDRESS
: Address to bind the metrics server, defaults to `0.0.0.0:9505`

DOCKERHUB_EXPORTER_WEB_PATH
: Path to bind the metrics server, defaults to `/metrics`

DOCKERHUB_EXPORTER_WEB_TIMEOUT
: Server metrics endpoint timeout, defaults to `10s`

DOCKERHUB_EXPORTER_REQUEST_TIMEOUT
: Timeout requesting DockerHub API, defaults to `5s`

DOCKERHUB_EXPORTER_USERNAME
: Username for the DockerHub authentication

DOCKERHUB_EXPORTER_PASSWORD
: Password for the DockerHub authentication

DOCKERHUB_EXPORTER_ORG
: Organizations to scrape metrics from, comma-separated list

DOCKERHUB_EXPORTER_REPO
: Repositories to scrape metrics from, comma-separated list

DOCKERHUB_EXPORTER_COLLECTOR_ORGS
: Enable collector for orgs, defaults to `true`

DOCKERHUB_EXPORTER_COLLECTOR_REPOS
: Enable collector for repos, defaults to `true`
