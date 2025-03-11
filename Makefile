postgres17:
	docker run --name postgres17-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

.PHONY: postgres