#!/usr/bin/env sh

set -x

migrate \
  -path /app/migrations \
  -database postgres://$APP_POSTGRES_USER:$APP_POSTGRES_PASSWORD@$APP_POSTGRES_HOST:$APP_POSTGRES_PORT/$APP_POSTGRES_DB?sslmode=disable \
  $*
