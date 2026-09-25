# Run all tests
test:
    go test ./...

# Run tests with race detection
test-race:
    go test -race ./...

# Run tests with coverage
test-cover:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

# Build the project
build:
    go build

# Run all benchmarks
bench:
    go test -bench=. -benchmem -benchtime=2s ./figures/

# Run benchmarks for a specific name (usage: just bench-run King)
bench-run name:
    go test -bench={{name}} -benchmem -benchtime=2s ./figures/
