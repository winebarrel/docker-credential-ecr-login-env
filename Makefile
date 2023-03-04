all: vet test build

.PHONY: build
build:
	go build ./cmd/docker-credential-ecr-login-with-env

.PHONY: vet
vet:
	go vet ./...

mock-helper:
	go build internal/test/mock-helper.go
bck-i-search: go bui_

.PHONY: test
test: mock-helper
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
