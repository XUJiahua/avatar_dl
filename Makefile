dev:
	go run .
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . && mkdir -p bin && mv avatar_dl bin/
