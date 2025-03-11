postgres17:
	docker run --name postgres17-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17-alpine createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres17-alpine dropdb simple_bank

.PHONY: postgres createdb dropdb