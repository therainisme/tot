#!/bin/bash

SERVICE_BINARY="/usr/local/bin/github-avatar-proxy"
SERVICE_FILE="/etc/systemd/system/github-avatar-proxy.service"
rm -f ${SERVICE_FILE} ${SERVICE_BINARY}

systemctl daemon-reload