# Changelog


## 1.1.0

The following sections list the changes for 1.1.0.

## Summary

 * Chg #135: Read secrets form files
 * Chg #135: Integrate standard web config
 * Enh #135: Integrate option pprof profiling

## Details

 * Change #135: Read secrets form files

   We have added proper support to load secrets like the password from files or
   from base64-encoded strings. Just provide the flags or environment variables for
   token or private key with a DSN formatted string like `file://path/to/file` or
   `base64://Zm9vYmFy`.

   https://github.com/promhippie/dockerhub_exporter/pull/135

 * Change #135: Integrate standard web config

   We integrated the new web config from the Prometheus toolkit which provides a
   configuration for TLS support and also some basic builtin authentication. For
   the detailed configuration you can check out the documentation.

   https://github.com/promhippie/dockerhub_exporter/pull/135

 * Enhancement #135: Integrate option pprof profiling

   We have added an option to enable a pprof endpoint for proper profiling support
   with the help of tools like Parca. The endpoint `/debug/pprof` can now
   optionally be enabled to get the profiling details for catching potential memory
   leaks.

   https://github.com/promhippie/dockerhub_exporter/pull/135


## 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #11: Refactor build tools and project structure

## Details

 * Change #11: Refactor build tools and project structure

   To have a unified project structure and build tooling we have integrated the
   same structure we already got within our GitHub exporter.

   https://github.com/promhippie/dockerhub_exporter/issues/11


## 0.2.0

The following sections list the changes for 0.2.0.

## Summary

 * Chg #2: Better naming for standard metrics
 * Chg #3: Enforce org or repo flag

## Details

 * Change #2: Better naming for standard metrics

   We named the metric to check if the exporter is working corrtly
   `dockerhub_valid_response` which doesn't reflect the Prometheus standards, so we
   renamed it to `dockerhub_up`.

   https://github.com/promhippie/dockerhub_exporter/issues/2

 * Change #3: Enforce org or repo flag

   We added some validation that makes sure that we got at least an organization or
   a repository defined to make sure we always got metrics to gather.

   https://github.com/promhippie/dockerhub_exporter/issues/3


## 0.1.0

The following sections list the changes for 0.1.0.

## Summary

 * Chg #10: Initial release of basic version

## Details

 * Change #10: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/promhippie/dockerhub_exporter/issues/10


