#!/bin/sh
set -e

DATABASE_HOST="database"
MIGRATIONS_PATH="/run/internal/database/migrations/mysql"
DATA_SOURCE_NAME="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${DATABASE_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"
OPTIONS="charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"

migrate -path="$MIGRATIONS_PATH" -database "${DATA_SOURCE_NAME}?${OPTIONS}" -verbose up

exec "$@"
