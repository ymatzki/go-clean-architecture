gen:
	wire ./...

mockgen:
	go generate ./...

start:
	go run main.go

setup:
	go install github.com/golang/mock/mockgen@v1.5.0
	go install github.com/google/wire/cmd/wire@latest

test:
	go test ./...
