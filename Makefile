build:
	@echo "Building the expense-cli application..."
	@go build ./cmd/expense-cli

test:
	@echo "Running tests..."
	@go test -v ./...

coverage:
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"