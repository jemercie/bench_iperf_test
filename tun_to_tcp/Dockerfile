FROM    debian 
# TODO: add debian version

COPY ./src/cli_unix_bin /bin/cli
COPY ./cli_as_daemon.conf /etc/init.d/tunnel-client


RUN apt update
RUN apt install -y iperf3 iproute2 openrc golang sudo htop
RUN chmod +x /bin/cli
RUN chmod +x /etc/init.d/tunnel-client
RUN sudo rc-update add tunnel-client default

ENTRYPOINT [ "sh", "-c", "sudo service tunnel-client start && sh"] 

# iperf3 -c 192.168.10.1 -B 192.168.11.1 -t 30
