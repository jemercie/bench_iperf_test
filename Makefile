
gen_bin:
	GOOS=linux GOARCH=amd64 go build -C ./src/cli -o ./bin/cli
	GOOS=linux GOARCH=amd64 go build -C ./src/srv -o ./bin/srv

.PHONY: gen_bin