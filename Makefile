GOCMD=go
BINARY_NAME=simple-login
BINARY_UNIX=$(BINARY_NAME)-unix

all: test build
build:
	$(GOCMD) build -o $(BINARY_NAME) -v main.go
run:
	make build
	./$(BINARY_NAME)
mock:
	mockery --all --keeptree
test:
	$(GOCMD) test ./... -cover -race -count=1 -coverprofile="coverage.out"
test-no-race:
	$(GOCMD) ./... -cover -count=1 -coverprofile="coverage.out"
lint:
	staticcheck ./...
tool:
	$(GOCMD) install github.com/vektra/mockery/v2@latest

	$(GOCMD) get honnef.co/go/tools/cmd/staticcheck@latest
build-linux:
	CGO_ENABLED=0 GOOS=linux $(GOCMD) build -o $(BINARY_UNIX) -a -installsuffix cgo -v main.go