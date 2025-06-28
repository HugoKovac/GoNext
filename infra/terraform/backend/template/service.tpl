[Unit]
Description=${description}
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=5
User=${user}
ExecStart=/usr/local/bin/${binary_name}
WorkingDirectory=/home/${user}

EnvironmentFile=${env_file}

# Security settings
NoNewPrivileges=yes
ProtectSystem=strict
ReadWritePaths=/tmp /var/log
PrivateTmp=yes

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=${binary_name}

[Install]
WantedBy=multi-user.target
