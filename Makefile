gen:
	go generate ./...
	wire ./...

start:
	go run main.go

setup:
	go install github.com/golang/mock/mockgen@v1.5.0
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

test:
	go test ./...
