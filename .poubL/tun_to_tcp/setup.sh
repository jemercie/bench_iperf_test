#!bin/bash

sudo cp ./src/cli_unix_bin /bin/cli
sudo cp ./cli_as_daemon.conf /etc/systemd/system/tunnel-client.service

sudo apt update
sudo apt install -y iperf3 inetutils-ping iproute2 openrc golang sudo htop net-tools 
sudo chmod +x /bin/cli
sudo chmod +x /etc/init.d/tunnel-client
sudo rc-update add tunnel-client default
