DB_URL=postgresql://root:secret@localhost:5432/shop?sslmode=disable

postgres:
	docker run --name postgres_shop -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres_shop createdb --username=root --owner=root shop

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down