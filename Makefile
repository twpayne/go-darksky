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

.PHONY: testdata
testdata:
	rm -f dstest/*.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -time 1556668800 > dstest/santamonica_20190501.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -extend hourly -units si > dstest/santamonica_hourly_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -exclude alerts,currently,daily,flags,minutely -units si > dstest/santamonica_exclude_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -exclude alerts,currently,daily,flags,minutely -extend hourly -units si > dstest/santamonica_exclude_hourly_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -lang fr > dstest/santamonica_fr.gen.go
