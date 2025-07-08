#!/bin/sh
set -e

chown -R dockerhub-exporter:dockerhub-exporter /var/lib/dockerhub-exporter
chmod 750 /var/lib/dockerhub-exporter

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet dockerhub-exporter.service; then
        systemctl restart dockerhub-exporter.service
    fi
fi
