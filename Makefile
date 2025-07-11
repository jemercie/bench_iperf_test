
bin:
	GOOS=linux GOARCH=amd64 go build -C ./cli -o cli
	GOOS=linux GOARCH=amd64 go build -C ./srv -o srv

.PHONY: bin