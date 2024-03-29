### Create the log files
```
mkdir -p /var/log/payments
touch /var/log/payments/info.log
touch /var/log/payments/error.log
chown root:root /var/log/payments
chmod 755 /var/log/payments
```

### Add the service file
/etc/systemd/system/payments.service

```
[Unit]
Description=payments.app
ConditionPathExists=/usr/local/bin
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
User=root
Group=root

WorkingDirectory=/usr/share/nginx/payments
ExecStart=/bin/bash -c "/usr/local/bin/payments --config=/usr/local/etc/payments.d 1>>/var/log/payments/info.log 2>>/var/log/payments/error.log"

Restart=on-failure
RestartSec=2

ExecStartPre=/bin/mkdir -p /var/log/payments
ExecStartPre=/bin/chown root:root /var/log/payments
ExecStartPre=/bin/chmod 755 /usr/local/bin/payments

StandardOutput=append:/var/log/payments/info.log
StandardError=append:/var/log/payments/error.log

[Install]
WantedBy=multi-user.target
```

### Enable the service
```
systemctl enable payments
```

### Update logrotate
/etc/logrotate.d/payments

```
/var/log/payments/*.log {
    size 1M
    create 0640 root root
    daily
    rotate 10
    missingok
    notifempty
    compress
    delaycompress
    sharedscripts
    postrotate
        /bin/kill -HUP `pgrep rsyslog 2> /dev/null` 2> /dev/null || true
    endscript
}
```

### Start the service
```
systemctl start payments
```

### Check the status
```
systemctl status payments
```