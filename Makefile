APP_ENVIRONMENT=local
DB_NAME=cashlog
TEST_DB_NAME=cashlog_test
DB_PORT=5432
TEST_DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=

run-monolith:
	APP_ENV=${APP_ENVIRONMENT} go run cmd/cashlog/main.go

test:
	go test ./...

db-setup: db-create db-migrate

db-create:
	createdb -p ${DB_PORT} -Opostgres -Eutf8 ${DB_NAME}

db-drop:
	dropdb -p ${DB_PORT} --if-exists -Upostgres ${DB_NAME}

db-migrate:
	migrate -database 'postgres://postgres@localhost:5432/${DB_NAME}?sslmode=disable&search_path=public' -path pkg/db/migrations up

db-migrate-down:
	migrate -database 'postgres://postgres@localhost:5432/${DB_NAME}?sslmode=disable&search_path=public' -path pkg/db/migrations down

db-test-setup: db-test-create

db-test-create:
	createdb -p ${TEST_DB_PORT} -Opostgres -Eutf8 ${TEST_DB_NAME}

db-test-drop:
	dropdb -p ${TEST_DB_PORT} --if-exists -Upostgres ${TEST_DB_NAME}
