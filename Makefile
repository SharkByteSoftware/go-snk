BASE_DIR := .
COV_PROFILE := dist/covprofile.out
TEST_OPTS := -vet=all -covermode=atomic -coverprofile=$(COV_PROFILE)

BENCH_PKGS := \
	slicex \
	mapx \
	containers/sets \
	containers/lists

BENCHMARKS := $(BENCH_PKGS:%=bench/%)
BENCH_COUNT := 1

.PHONY: default all test bench bench/% vet lint tidy clean dist godoc browse update-pkg-go-dev

default: test lint lint-tests

all: default bench

test: | dist
	$(printTarget)
	go test $(TEST_OPTS) ./...
	go tool cover --html=$(COV_PROFILE) -o dist/coverage_report.html

bench: $(BENCHMARKS)
bench/%:
	$(printTarget)
	cd $* && go test -benchmem -count $(BENCH_COUNT) -bench .

lint:
	$(printTarget)
	@golangci-lint run

lint-tests:
	$(printTarget)
	@golangci-lint run --tests -c .golangci-tests.yml

tidy:
	$(printTarget)
	go mod tidy -v

clean:
	@go clean ./...
	@go clean -testcache
	@rm -rf dist/

dist:
	$(printTarget)
	@mkdir -p $(@)

godoc:
	$(printTarget)
	@go install golang.org/x/tools/cmd/godoc@latest
	godoc -v -http=:6060

browse:
	open http://localhost:6060/pkg/github.com/SharkByteSoftware/go-snk


update-pkg-go-dev:
	$(printTarget)
	$(eval LATEST_TAG := $(shell git ls-remote --tags origin | grep -E 'v[0-9]+\.[0-9]+\.[0-9]+$$' | awk '{print $$2}' | sed 's#refs/tags/##' | sort -V | tail -n 1))
	@test -n "$(LATEST_TAG)" || (echo "Could not determine latest tag" && exit 1)
	@echo "Refreshing pkg.go.dev with: $(LATEST_TAG)"
	@curl -s "https://proxy.golang.org/github.com/!shark!byte!software/go-snk/@v/$(LATEST_TAG).info" | jq

#Helper function to pretty print targets as they execute
TARGET_COLOR := \033[0;32m
NO_COLOR := \033[m
CURRENT_TARGET = $(@)

define printTarget
	@printf "%b" "\n$(TARGET_COLOR)$(CURRENT_TARGET):$(NO_COLOR) $^\n";
endef