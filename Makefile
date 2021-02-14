DB_NAME=cashlog
TEST_DB_NAME=cashlog_test
DB_PORT=5432
TEST_DB_PORT=5432

db-setup: db-create

db-create:
	createdb -p ${DB_PORT} -Opostgres -Eutf8 ${DB_NAME}

db-drop:
	dropdb -p ${DB_PORT} --if-exists -Upostgres ${DB_NAME}

db-test-setup: db-test-create

db-test-create:
	createdb -p ${TEST_DB_PORT} -Opostgres -Eutf8 ${TEST_DB_NAME}

db-test-drop:
	dropdb -p ${TEST_DB_PORT} --if-exists -Upostgres ${TEST_DB_NAME}
