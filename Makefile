BASE_DIR = .
COV_PROFILE = dist/covprofile.out
TEST_OPTS = -vet=all -cover -covermode=atomic -coverprofile=$(COV_PROFILE)

BENCH_PKGS = \
	slices \
	maps \
	ds
BENCHMARKS = $(BENCH_PKGS:%=bench/%)
BENCH_COUNT = 1

default: test vet

all: test vet

test: | dist
	$(printTarget)
	go test $(TEST_OPTS) ./...
	go tool cover --html=$(COV_PROFILE) -o dist/coverage_report.html

bench: $(BENCHMARKS)
bench/%:
	$(printTarget)
	cd $(@F) && go test -benchmem -count $(BENCH_COUNT) -bench .

vet:
	$(printTarget)
	@golangci-lint run

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

#Helper function to pretty print targets as they execute
TARGET_COLOR := \033[0;32m
NO_COLOR := \033[m
CURRENT_TARGET = $(@)

define printTarget
	@printf "%b" "\n$(TARGET_COLOR)$(CURRENT_TARGET):$(NO_COLOR) $^\n";
endef