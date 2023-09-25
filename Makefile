MAIN_ENTRY="cmd/gocalc/main.go"
COVERAGE_FILE="coverage.out"

build:
	go build $(MAIN_ENTRY)

run:
	go run $(MAIN_ENTRY)

format:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v -coverprofile=$(COVERAGE_FILE)

test-coverage:
	go tool cover -html=$(COVERAGE_FILE)