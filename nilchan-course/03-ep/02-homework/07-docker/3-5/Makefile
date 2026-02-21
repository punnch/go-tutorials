include .env
export

# docker
deploy-app:
	docker compose up -d application

deploy-postgres:
	docker compose up -d postgres

# migrate
migrate-up:
	migrate -path migrations -database ${DB_URL} up

migrate-down:
	migrate -path migrations -database ${DB_URL} down

# other
service-run:
	go run main.go