test:
	go test -v -cover ./...

createdb:
	docker exec -it postgresql-docker createdb -U postgres simple_bank

dropdb:
	docker exec -it postgresql-docker dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration/postgres -database ${DATABASE} -verbose up

migrateup_test:
	migrate -path db/migration/postgres -database "postgresql://postgres:kokoro@localhost:5432/simple_bank_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/postgres -database ${DATABASE} -verbose down

migratedown_test:
	migrate -path db/migration/postgres -database "postgresql://postgres:kokoro@localhost:5432/simple_bank_test?sslmode=disable" -verbose down

server:
	go run main.go

.PHONY: test createdb dropdb migrateup migrateup_test migratedown migratedown_test
