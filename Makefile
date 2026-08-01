.PHONY: run test build

run:
	GOFLAGS="-buildvcs=false" swag init -g cmd/api/main.go -d ./ --parseDependency
	go run cmd/api/main.go

test:
	go test -v ./...

build:
	go build -o bin/api cmd/api/main.go
	
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate-up:
	migrate -path shared/database/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path shared/database/migrations -database "$(DATABASE_URL)" -verbose down