.PHONY: build test clean

build:
	echo "Building..."
	go build -o bin/zambone ./cmd/zambone

test:
	echo "Testing..."
	go test ./... -v

clean:
	echo "Cleaning binaries..."
	rm -fr bin/