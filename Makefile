.PHONY: server fmt vet cyclo

SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | grep -v /gen |cut -d "/" -f2 | uniq)

run:
	go run main.go

server:
	go run main.go server

fmt:
	go get -u golang.org/x/tools/cmd/goimports
	@echo "Reformatting..."
	@goimports -l -w $(SOURCE_DIRS)

vet:
	@echo "Running vet..."
	@go vet ./...

cyclo:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo "Running cyclo..."
	@gocyclo -over 7 $(SOURCE_DIRS)