# VARIABLE REFERENCE:
#
# Test-specific variables:
#   T=<pattern>       - Test name pattern (e.g., VarName/simple)
#   COV=true          - Enable coverage
#
#
# Examples:
#   make test T=VarName/simple                            # Run only the test with the given name
#   make test COV=true                                    # Run all unit tests with coverage

COVERAGEARGS?=-race -coverprofile=coverage.txt -covermode=atomic

ifneq ($(origin T), undefined)
	RUNARGS = -run='$(T)'
endif

ifneq ($(origin COV), undefined)
	RUNARGS += $(COVERAGEARGS)
endif

tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0

fmt:
	@echo "==> Fixing source code formatting..."
	golangci-lint fmt ./src/...

lint:
	@echo "==> Checking source code against linters and fixing..."
	golangci-lint run --fix ./src/...

lintcheck:
	@echo "==> Checking source code against linters and fixing..."
	golangci-lint run ./src/...

test:
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Running unit tests on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	go test ./src/... \
		-timeout=30s \
		-parallel=4 \
		-v \
		$(RUNARGS) $(TESTARGS) \
		-count 1;

.PHONY: tools fmt lint lintcheck test
