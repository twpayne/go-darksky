.PHONY: none
none:

.PHONY: coverage.out
coverage.out:
	go test -cover -coverprofile=$@ ./... || ( rm -f $@ ; false )

.PHONY: format
format:
	find . -name \*.go | xargs gofumports -w

.PHONY: html-coverage
html-coverage: coverage.out
	go tool cover -html=$<

.PHONY: install-tools
install-tools:
	GO111MODULE=off go get -u \
		golang.org/x/tools/cmd/cover \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		mvdan.cc/gofumpt/gofumports

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

.PHONY: test
test:
	go test ./...
