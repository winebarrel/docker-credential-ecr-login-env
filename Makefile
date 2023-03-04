all: vet test build

.PHONY: build
build:
	go build ./cmd/docker-credential-ecr-login-env

.PHONY: vet
vet:
	go vet ./...

mock-helper: internal/test/mock-helper.go
	go build internal/test/mock-helper.go

.PHONY: test
test: mock-helper
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
