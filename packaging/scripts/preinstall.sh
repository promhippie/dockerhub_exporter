#!/bin/sh
set -e

if ! getent group dockerhub-exporter >/dev/null 2>&1; then
    groupadd --system dockerhub-exporter
fi

if ! getent passwd dockerhub-exporter >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/dockerhub-exporter --shell /bin/bash -g dockerhub-exporter dockerhub-exporter
fi
