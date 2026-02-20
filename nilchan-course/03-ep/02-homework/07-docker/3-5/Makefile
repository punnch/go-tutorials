include .env
export

run-http-app:
	@docker run -p 8081:8081 http-app-image

migrate-up:
	migrate -path migrations -database ${DB_URL} up