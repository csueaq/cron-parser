.PHONY=build test

GO111MODULE?=on
GO_BIN?=app
GO?=go

.EXPORT_ALL_VARIABLES:

test:
	$(GO) test ./... -cover -coverprofile coverage.out

build: test
	$(GO) build -o $(GO_BIN) cmd/main.go
