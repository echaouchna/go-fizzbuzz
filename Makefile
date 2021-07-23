NAME=go-fizzbuzz

run: .mod-download
	go run .

build: .mod-download
	mkdir -p bin
	go build -v --ldflags '-s -extldflags "-static"' -o bin/${NAME} .

build-docker:
	docker build --no-cache -t echaouchna/${NAME} .

test: .mod-download
	go test

.mod-download:
	go mod download

.PHONY: clean

clean:
	-rm -rf bin