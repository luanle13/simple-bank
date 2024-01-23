env:
	docker-compose up -d

install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=root simple_bank

migrateup:
	migrate -path database/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: env createdb dropdb migrateup migratedown sqlc