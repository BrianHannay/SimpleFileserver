all: bin/fileserver

run: bin/fileserver
	bin/fileserver

bin/fileserver: fileserver.go go format
	go build -o bin/fileserver fileserver.go

go:
	go version

format:
	gofmt -w -e -s -d $$(find . -type f -name '*.go')

clean:
	rm -r bin/*

.PHONY: all run go format clean
