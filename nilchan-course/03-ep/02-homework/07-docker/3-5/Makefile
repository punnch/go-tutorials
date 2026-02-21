include .env
export

run-http-app:
	docker run -p 8081:8081 http-app-image

migrate-up:
	migrate -path migrations -database ${DB_URL} up

migrate-down:
	migrate -path migrations -database ${DB_URL} down

service-run:
	go run main.go

postgres-up:
	docker run -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p 5432:5432 -v ./pgdata:/var/lib/postgresql -d postgres:18-bookworm