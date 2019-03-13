
.PHONY: all test verbose

all: test

verbose:
	go test -v ./...

test:
	go test ./...

