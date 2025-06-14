.PHONY: build test clean run

# Build the application
build:
	go build -o bin/k8s-gpu-analyzer ./cmd/k8s-gpu-analyzer

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Run the application
run: build
	./bin/k8s-gpu-analyzer

# Install dependencies
deps:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run all checks
check: fmt vet test

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/k8s-gpu-analyzer-linux-amd64 ./cmd/k8s-gpu-analyzer
	GOOS=darwin GOARCH=amd64 go build -o bin/k8s-gpu-analyzer-darwin-amd64 ./cmd/k8s-gpu-analyzer
	GOOS=darwin GOARCH=arm64 go build -o bin/k8s-gpu-analyzer-darwin-arm64 ./cmd/k8s-gpu-analyzer
	GOOS=windows GOARCH=amd64 go build -o bin/k8s-gpu-analyzer-windows-amd64.exe ./cmd/k8s-gpu-analyzer
