
**build:**
``` sh
cd benchmark_iperf_test/
make build
```

**start server:**
```
cd tcp_to_tun/
docker run -it --cap-add=NET_ADMIN --device=/dev/net/tun --network tcp-tunnel --rm --name server srv
```

**start client:**
```
cd tun_to_tcp/
docker run -it --cap-add=NET_ADMIN --device=/dev/net/tun --network tcp-tunnel --rm --name client cli
```


- `iperf3 -s -B 192.168.10.1`
- `iperf3 -c 192.168.10.1 -B 192.168.11.1`
-  and launch `tun_to_tcp` and `tcp_tp_tun`

`ifconfig`
`netstat -rn`
`netstat -an | grep 5201`
`tcpdump -i utun5 -v` -- utun5 or the created output tun
`sudo tcpdump -i utun5 -n -S tcp`
`sudo tcpdump -i utun4 -n -S tcp`

ping server

print routes:
`netstat -rn -f inet`
delete route:
`route delete <destination ip or interface>`

[disable firewal stealth mode and now kernel answer to pings](https://discussions.apple.com/thread/2639727?sortBy=rank)

and damn it doesn't work krkrkr

## iperf3

[documentation](https://iperf.fr/iperf-doc.php)

`--logfile </path/name>` to get the output in a file


## Docker

[compose capp_add?](https://forums.docker.com/t/docker-compose-order-of-cap-drop-and-cap-add/97136)

[docker network types](https://devopssec.fr/article/fonctionnement-manipulation-reseau-docker)
-> default network is type `bridge`

[docker compose connection between containers](https://stackoverflow.com/questions/65042615/docker-compose-connection-between-containers)



## Dagger

`dagger init --name=tcp_tunnel --sdk=go` to init the dagger module

clean the cache:
>`dagger core engine local-cache prune`

dagger option to print the full trace:
>`dagger call <function> <options> --progress=plain`

[`WithExposedPort(<port>, dagger.ContainerWithExposedPortOpts{Option:status})` documentation](https://docs.dagger.io/reference/typescript/api/client.gen/type-aliases/ContainerWithExposedPortOpts)

[apt update is mandatory](https://askubuntu.com/questions/337198/is-sudo-apt-get-update-mandatory-before-every-package-installation)

### Cross Compile in GO
to [cross compile in go](https://freshman.tech/snippets/go/cross-compile-go-programs/) we just have to change env variables like this:
>`GOOS=linux GOARCH=arm64 go build -o cli_unix_bin`
>`GOOS=linux GOARCH=arm64 go build -o srv_unix_bin`

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

### OpenRC
[documentation to write openrc script receipe](https://www.funtoo.org/Openrc)

### Env to test config

#### SRV

`dagger`
```sh
container | from ubuntu | with-mounted-file /bin/srv ./tcp_to_tun/srv_unix_bin | with-mounted-file /etc/init.d/monitor_srv ./.dagger/srv_as_daemon.conf | with-exposed-port 4663 | terminal
```
```sh
apt update
mkdir -p /dev/net
touch /dev/net/tun
apt install golang -y
apt install -y iperf3 iproute2 openrc
chmod +x /bin/srv
chmod +x etc/init.d/monitor_srv
openrc default
touch /run/openrc/softlevel
rc-update add monitor_srv default
service monitor_srv start

```


#### CLI

`dagger`
```sh
container | from ubuntu | with-mounted-file /bin/cli ./tun_to_tcp/cli_unix_bin | with-mounted-file /etc/init.d/monitor_cli ./.dagger/cli_as_daemon.conf | with-exposed-port 4663 | terminal
```
```sh
apt update
mkdir -p /dev/net
apt install golang -y
touch /dev/net/tun
apt install -y iperf3 iproute2 openrc
chmod +x /bin/cli
chmod +x etc/init.d/monitor_cli
openrc default
touch /run/openrc/softlevel
rc-update add monitor_cli default
service monitor_cli start

```
