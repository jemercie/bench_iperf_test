FROM    debian 
# TODO: add debian version

COPY ./src/srv_unix_bin /bin/srv
COPY ./srv_as_daemon.conf /etc/init.d/tunnel-server

RUN apt update
RUN apt install -y iperf3 iproute2 openrc golang sudo htop
RUN chmod +x /bin/srv
RUN chmod +x /etc/init.d/tunnel-server
RUN sudo rc-update add tunnel-server default


ENTRYPOINT ["sh", "-c", "sudo service tunnel-server start && sh"]
# iperf3 -s -B 192.168.10.1
# ENTRYPOINT ["sh", "-c", "iperf3 -s -D --logfile /var/log/iperf3_cli_log -B 192.168.10.1 && sudo service tunnel-server start && sh"]
