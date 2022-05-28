# Develop
gen:
	go generate ./...
	wire ./...

setup:
	go install github.com/golang/mock/mockgen@v1.5.0
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

test:
	go test ./...

# Database
init-db:
	docker-compose up -d

clean-db:
	docker-compose down

# Run Application
start: init-db
	go run main.go
