test:
	echo "Happy Lucky Smile Ya!"

createdb:
	docker exec -it postgresql-docker createdb -U postgres simple_bank


dropdb:
	docker exec -it postgresql-docker dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migrations/postgres -database ${DATABASE} -verbose up

migratedown:
	migrate -path db/migrations/postgres -database ${DATABASE} -verbose down

.PHONY: test migrateup migratedown
