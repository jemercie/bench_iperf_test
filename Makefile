
bin:
	GOOS=linux GOARCH=amd64 go build -C="./src/cli" -o cli
	GOOS=linux GOARCH=amd64 go build -C="./src/srv" -o srv

.PHONY: bin