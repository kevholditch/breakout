default: build

.PHONY: build
build:
	@env GOMODULE111=on find ./cmd/* -maxdepth 1 -type d -exec go build "{}" \;

.PHONY: test
test:
	@go test -v ./...

.PHONY: vet
vet:
	@go vet -v ./...

.PHONY: check
check: vet