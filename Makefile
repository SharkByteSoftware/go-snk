NAME = go-sink
BASE_DIR = .
TEST_OPTS =-vet=all

default: test vet

all: test vet

test:
	$(printTarget)
	go test $(TEST_OPTS) ./...

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