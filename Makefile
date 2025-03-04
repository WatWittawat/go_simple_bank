postgres:
	docker compose up -d

createdb:
	docker exec -it postgres createdb --username=myuser --owner=myuser simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup_1:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown_1:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/WatWittawat/go_simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server migratedown_1 migrateup_1
