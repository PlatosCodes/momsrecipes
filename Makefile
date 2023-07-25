DB_URL=postgresql://root:bluecomet@localhost:5432/momsrecipes?sslmode=disable
network:
	docker network create momsrecipes-network
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=bluecomet -d postgres:15-alpine
	
createdb: 
	docker exec -it postgres createdb --username=root --owner=root momsrecipes

dropdb:
	docker exec -it postgres dropdb --username=root momsrecipes

citext:
	docker exec -it postgres psql --username=root momsrecipes -c "CREATE EXTENSION IF NOT EXISTS citext;"


migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedown: 
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1: 
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/database.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/PlatosCodes/momsrecipes/db/sqlc Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto	

evans: 
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb citext migrateup migrateup1 migratedown migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans