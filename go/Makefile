default: help

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: test
test:
	@go test -v ./...

.PHONY: build
build:
	@go build -o bin/main.go

.PHONY: clean
clean:
	@go clean
	rm -rf bin/main.go

.PHONY: tidy
	@go mod tidy

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: help
help:
	@echo ""
	@echo "Usage:"
	@echo "make <command>"
	@echo "You can replace <command> by test, run, build, clean, tidy, fmt"
