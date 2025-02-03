DB_URL=postgresql://deng:deng@192.168.193.158:5432/simple_bank?sslmode=disable

postgres:
	docker run --name simpleBank --network bank-network -e POSTGRES_USER=deng -e POSTGRES_PASSWORD=deng -p 5432:5432 -d postgres:17-alpine

createdb:
	docker exec -it simpleBank createdb --username=deng --owner=deng simple_bank

dropdb:
	docker exec -it simpleBank dropdb -U deng simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown1 migratedown migratedown1 sqlc test server mock
