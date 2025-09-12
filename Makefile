NAME = go-sink
BASE_DIR = .
TEST_OPTS =-vet=all

BENCH_PKGS = \
	slices \
	maps \
	sets
BENCHMARKS = $(BENCH_PKGS:%=bench/%)
BENCH_COUNT = 1

default: test vet

all: test vet

test:
	$(printTarget)
	go test $(TEST_OPTS) ./...

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

#Helper function to pretty print targets as they execute
TARGET_COLOR := \033[0;32m
NO_COLOR := \033[m
CURRENT_TARGET = $(@)

define printTarget
	@printf "%b" "\n$(TARGET_COLOR)$(CURRENT_TARGET):$(NO_COLOR) $^\n";
endef