postgres-init:
	docker run \
	--name postgres15 \
	-p 5433:5432 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=password \
	-d \
	postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

create-db:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

drop-db:
	docker exec -it postgres15 dropdb go-chat

migrate-up:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgres-init postgres create-db drop-db migrate-up migrate-down
