## builder
FROM golang:1.14 as builder

WORKDIR /app
COPY . .
RUN make build

## app wrapper
FROM ubuntu:focal

WORKDIR /app
COPY --from=builder /app/build ./build/

ENV APP_POSTGRES_DB=app
ENV APP_POSTGRES_USER=app
ENV APP_POSTGRES_PASSWORD=app_pass
ENV APP_POSTGRES_PORT=5432

ENV APP_REDIS_DB=0
ENV APP_REDIS_PORT=6379

EXPOSE 80

CMD ["/app/build/web"]

