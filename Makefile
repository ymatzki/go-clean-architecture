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
up-db:
	docker-compose up -d

down-db:
	docker-compose down

clean-db:
	docker volume rm go-clean-architecture_data

init-db: down-db clean-db up-db

# Run Application
start: init-db
	go run main.go
