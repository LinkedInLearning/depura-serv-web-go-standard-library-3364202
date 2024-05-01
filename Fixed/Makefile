DB_PASSWORD=postgres
DB_USER=postgres
DB_NAME=postgres

run-db:
	DB_NAME=$(DB_NAME) DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) \
		docker compose up -d

stop-db:
	docker compose down --remove-orphans --volumes

run-app: run-db
	DB_NAME=$(DB_NAME) DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) \
		go run main.go
