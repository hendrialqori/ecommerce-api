go:
	go run cmd/web/main.go

include .env

DATABASE_URL="mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

migrate.create:
	migrate create -ext sql -dir db/migrations $(name)

migrate.up:
	migrate -database $(DATABASE_URL) -path db/migrations up

migrate.down:
	migrate -database $(DATABASE_URL) -path db/migrations down

migrate.force:
	migrate -database $(DATABASE_URL) -path db/migrations force $(version)