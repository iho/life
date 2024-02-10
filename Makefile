default: lint test run

lint:
	golangci-lint run

test:
	go test -v -cover ./...

run:
	go run main.go