.PHONY:

test:
	go test -v ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...