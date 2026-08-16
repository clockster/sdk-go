.PHONY: spec generate operations lint test check

spec:
	go run ./internal/cmd/syncspec

generate:
	go run ./internal/cmd/generate

operations:
	go run ./internal/cmd/checkops

lint:
	test -z "$$(gofmt -l .)" || { gofmt -l .; echo 'Run gofmt -w .'; exit 1; }
	go vet ./...

test:
	go test -race ./...

check: lint test operations
