NAME = sb
BASE_DIR = .
TEST_OPTS =-vet=all

default: test vet

all: test vet

test:
	@go test $(TEST_OPTS) ./...

vet:
	@golangci-lint run --output.text.path stdout

tidy:
	$(printTarget)
	go mod tidy -v

clean:
	@go clean ./...
	@go clean -testcache
	@rm -rf $(NAME) $(CREATE_DIRS)
