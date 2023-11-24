#!/bin/bash

SERVICE_PORT=5203
SERVICE_BINARY="/usr/local/bin/github-avatar-proxy"
SERVICE_FILE="/etc/systemd/system/github-avatar-proxy.service"
rm -f ${SERVICE_FILE} ${SERVICE_BINARY}

go build -o $SERVICE_BINARY
chmod +x $SERVICE_BINARY

cat > ${SERVICE_FILE} <<- EOF
[Unit]
Description=github-avatar-proxy systemd service unit file.

[Service]
ExecStart=/usr/local/bin/github-avatar-proxy -p ${SERVICE_PORT}

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
echo "systemctl enable --now github-avatar-proxy"
