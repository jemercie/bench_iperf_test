5 term

- `iperf3 -s -B 192.168.10.1`
- `iperf3 -c 192.168.10.1 -B 192.168.11.1`
-  and launch `tun_to_tcp` and `tcp_tp_tun`

`ifconfig`
`netstat -rn`
`netstat -an | grep 5201`
`tcpdump -i utun5 -v` -- utun5 or the created output tun
`sudo tcpdump -i utun5 -n -S tcp`
`sudo tcpdump -i utun4 -n -S tcp`

print routes:
`netstat -rn -f inet`
delete route:
`route delete <destination ip or interface>`

[disable firewal stealth mode and now kernel answer to pings](https://discussions.apple.com/thread/2639727?sortBy=rank)

-> the problem is that the syn, ack of the tcp handshake is never sent anywhere so problem.


and damn it doesn't work krkrkr

### Dagger

`dagger init --name=tcp_tunnel --sdk=go` to init the dagger module

clean the cache:
>`dagger core engine local-cache prune`

dagger option to print the full trace:
>`dagger call <function> <options> --progress=plain`

### Cross Compile in GO
to [cross compile in go](https://freshman.tech/snippets/go/cross-compile-go-programs/) we just have to change env variables like this:
>`GOOS=linux GOARCH=amd64 go build -o cli_unix_bin`
>`GOOS=linux GOARCH=amd64 go build -o srv_unix_bin`

### Need more than one entrypoint
[run cmd as a daemon/service on alpine](https://medium.com/@mfranzon/how-to-create-and-manage-a-service-in-an-alpine-linux-container-93a97d5dad80)

```sh

apk add openrc
chmod +x </path/to/my/cnd-pgm-or-other>
vim etc/init.d/monitor-thing

echo '
#!/sbin/openrc-run
name=""
description=""
command="<path/to/cmd>"
command_background=true
pidfile="/run/monitor_<thing>.pid"
'
chmod +x <etc/init.d/monitor-thing>

openrc default
rc-update add monitor_thing default
service monitor_thing start

```

### Env to test config

#### CLI

`dagger`

`dagger call build --cli_bin=./tcp_to_tun/cli_unix_bin --srv_bin=./tun_to_tcp/srv_unix_bin --cli_daemon_conf=./.dagger/cli_as_daemon.conf --srv_daemon_conf=./.dagger/srv_as_daemon.conf  --progress=plain`

```sh
apk add iperf3 iproute2 openrc
chmod +x /bin/cli
chmod +x etc/init.d/monitor_cli
openrc default
touch /run/openrc/softlevel
rc-update add monitor_cli default
service monitor_cli start

```

#### SRV

`dagger`

`dagger call build --cli_bin=./tcp_to_tun/cli_unix_bin --srv_bin=./tun_to_tcp/srv_unix_bin --cli_daemon_conf=./.dagger/cli_as_daemon.conf --srv_daemon_conf=./.dagger/srv_as_daemon.conf  --progress=plain`

```sh
apk add iperf3 iproute2 openrc
chmod +x /bin/srv
chmod +x etc/init.d/monitor_srv
openrc default
touch /run/openrc/softlevel
rc-update add monitor_srv default
service monitor_srv start

```

tester a la main, ctrl+z puis bg pour gettre en background damn
