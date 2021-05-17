Change: Better naming for standard metrics

We named the metric to check if the exporter is working corrtly
`dockerhub_valid_response` which doesn't reflect the Prometheus standards, so we
renamed it to `dockerhub_up`.

https://github.com/promhippie/dockerhub_exporter/issues/2
