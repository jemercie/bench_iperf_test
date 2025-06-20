

bin:
	GOOS=linux GOARCH=arm64 go build -C ./tcp_to_tun/src -o srv_unix_bin
	GOOS=linux GOARCH=arm64 go build -C ./tun_to_tcp/src -o cli_unix_bin

build: bin
	docker build -t srv ./tcp_to_tun/
	docker build -t cli ./tun_to_tcp/

run_cli:
	docker run -it --cap-add=NET_ADMIN --device=/dev/net/tun cli

run_srv:
	docker run -it --cap-add=NET_ADMIN --device=/dev/net/tun srv

run: run_srv run_cli

.PHONY: bin uild run_srv run_cli run
