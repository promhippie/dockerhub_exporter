# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #11: Refactor build tools and project structure

## Details

 * Change #11: Refactor build tools and project structure

   To have a unified project structure and build tooling we have integrated the same structure we
   already got within our GitHub exporter.

   https://github.com/promhippie/dockerhub_exporter/issues/11


# Changelog for 0.1.0

The following sections list the changes for 0.1.0.

## Summary

 * Chg #10: Initial release of basic version

## Details

 * Change #10: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/promhippie/dockerhub_exporter/issues/10


# Changelog for 0.2.0

The following sections list the changes for 0.2.0.

## Summary

 * Chg #2: Better naming for standard metrics
 * Chg #3: Enforce org or repo flag

## Details

 * Change #2: Better naming for standard metrics

   We named the metric to check if the exporter is working corrtly `dockerhub_valid_response`
   which doesn't reflect the Prometheus standards, so we renamed it to `dockerhub_up`.

   https://github.com/promhippie/dockerhub_exporter/issues/2

 * Change #3: Enforce org or repo flag

   We added some validation that makes sure that we got at least an organization or a repository
   defined to make sure we always got metrics to gather.

   https://github.com/promhippie/dockerhub_exporter/issues/3


