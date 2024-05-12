postgresinit:
	docker run --name postgres15 -d -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres postgres:15-alpine 

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root blogator

dropdb:
	docker exec -it postgres15 dropdb blogator

migrateupdb:
	cd sql/schema/ && goose postgres "postgres://root:postgres@localhost:5434/blogator" up

migratedowndb:
	cd sql/schema/ && goose postgres "postgres://root:postgres@localhost:5434/blogator" down

.PHONY: postgresinit postgres createdb dropdb