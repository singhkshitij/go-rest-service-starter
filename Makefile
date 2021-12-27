.PHONY: server build run fmt vet cyclo checks

SOURCE_DIRS=$(shell go list ./... | cut -d "/" -f4 | uniq)

run:
	go run main.go

server:
	go run main.go server

build:
	go build -o bin/superdms

fmt:
	go get -u golang.org/x/tools/cmd/goimports
	@echo "Reformatting..."
	@gofmt -s -w .
	@goimports -l -w $(SOURCE_DIRS)

vet:
	@echo "Running vet..."
	@go vet ./...

cyclo:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo "Running cyclo..."
	@gocyclo -over 5 $(SOURCE_DIRS)

print:
	echo $(SOURCE_DIRS)

checks: fmt vet cyclo