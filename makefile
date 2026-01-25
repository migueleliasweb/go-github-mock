
tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0

fmt:
	@echo "==> Fixing source code formatting..."
	golangci-lint fmt ./...

lint:
	@echo "==> Checking source code against linters and fixing..."
	golangci-lint run --fix ./...

.PHONY: tools fmt lint
