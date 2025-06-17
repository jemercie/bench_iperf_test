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

to [cross compile in go](https://freshman.tech/snippets/go/cross-compile-go-programs/) we just have to change env variables like this:
>`GOOS=linux GOARCH=amd64 go build -o cli_unix_bin`


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