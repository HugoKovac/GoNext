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

# Environment variables (uncomment and modify as needed)
# Environment=PORT=${port}
# Environment=LOG_LEVEL=info
# Environment=DATABASE_URL=your_database_url

# Security settings
NoNewPrivileges=yes
ProtectSystem=strict
ProtectHome=yes
ReadWritePaths=/tmp /var/log
PrivateTmp=yes

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=${binary_name}

[Install]
WantedBy=multi-user.target
