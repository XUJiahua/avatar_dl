dev:
	go run .
dev_help:
	go run . -h
clean:
	rm failure_uris.txt logrus.log && rm -rf download
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . && mkdir -p bin && mv avatar_dl bin/
build_darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build . && mkdir -p bin && mv avatar_dl bin/
