postgres:
	docker run --name simpleBank -e POSTGRES_USER=deng -e POSTGRES_PASSWORD=deng -p 5432:5432 -d postgres:17-alpine

createdb:
	docker exec -it simpleBank createdb --username=deng --owner=deng simple_bank

dropdb:
	docker exec -it simpleBank dropdb -U deng simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://deng:deng@192.168.193.158:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://deng:deng@192.168.193.158:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
