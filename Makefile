.PHONY=build

GO111MODULE?=on
GO_BIN?=app
GO?=go

.EXPORT_ALL_VARIABLES:

build:
	$(GO) build -a -installsuffix nocgo -o $(GO_BIN) cmd/main.go
