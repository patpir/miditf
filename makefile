
.PHONY: all test verbose coverage

all: test

verbose:
	go test -v ./...

test:
	go test ./...

coverage:
	go test --cover ./...

