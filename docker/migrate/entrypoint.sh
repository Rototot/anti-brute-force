#!/usr/bin/env sh

set -x

waitforit -host $APP_POSTGRES_HOST -port $APP_POSTGRES_PORT -timeout=15 -debug -- echo "db is up"

migrate \
  -path /app/migrations \
  -database postgres://$APP_POSTGRES_USER:$APP_POSTGRES_PASSWORD@$APP_POSTGRES_HOST:$APP_POSTGRES_PORT/$APP_POSTGRES_DB?sslmode=disable \
  $*
