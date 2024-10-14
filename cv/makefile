.PHONY: fmt vet lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

run: fmt vet
	go run cmd/main.go