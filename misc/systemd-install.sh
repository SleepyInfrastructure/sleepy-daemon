cp service-linux.sh /opt/sleepy-daemon/service-linux.sh
cp systemd.service /etc/systemd/system/sleepy-daemon.service
systemctl daemon-reload
# systemctl start sleepy-daemon.service
systemctl enable sleepy-daemon.service