
.PHONY: all test verbose coverage

all: test

verbose:
	go test -v ./...

test:
	go test ./...

coverage:
	go test --cover ./...

uncovered:
	go test -coverprofile coverage.txt ./...
	@echo
	@echo "Not covered by tests:" 
	@grep -E " 0$$" coverage.txt

