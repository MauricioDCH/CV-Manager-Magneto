.PHONY: fmt vet lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

run: fmt vet lint
	go run cmd/main.go
