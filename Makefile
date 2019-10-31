.PHONY: build clean

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/permissions main.go

clean:
	rm -rf ./bin
