#!/bin/sh
set -e

systemctl stop dockerhub-exporter.service || true
systemctl disable dockerhub-exporter.service || true
