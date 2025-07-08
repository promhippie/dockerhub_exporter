#!/bin/sh
set -e

if [ ! -d /var/lib/dockerhub-exporter ]; then
    userdel dockerhub-exporter 2>/dev/null || true
    groupdel dockerhub-exporter 2>/dev/null || true
fi
