
deploy-local:
	cd deploy/ && \
	docker-compose -p "green" up -d

run-local:
    SERVER_ADDRESS=localhost SERVER_PORT=3005 \
    DB_USER=postgres \
    DB_PASS=changeme \
    DB_HOST=localhost \
    DB_PORT=5432 \
    DB_NAME=green \
    go run ./cmd/green/main.go
