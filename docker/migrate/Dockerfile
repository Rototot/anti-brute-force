FROM alpine:latest as waitforit

RUN apk add curl
ENV WAITFORIT_VERSION="v2.4.1"
RUN curl -o /usr/local/bin/waitforit -sSL https://github.com/maxcnunes/waitforit/releases/download/$WAITFORIT_VERSION/waitforit-linux_amd64 && \
    chmod +x /usr/local/bin/waitforit

FROM migrate/migrate:latest

WORKDIR /app/migrations

COPY --from=waitforit /usr/local/bin/waitforit /usr/local/bin/waitforit
COPY ./docker/migrate/entrypoint.sh /opt/
COPY ./db/migrations /app/migrations

ENV APP_POSTGRES_HOST=app
ENV APP_POSTGRES_DB=app
ENV APP_POSTGRES_USER=app
ENV APP_POSTGRES_PASSWORD=app_pass
ENV APP_POSTGRES_PORT=5432

ENTRYPOINT ["/opt/entrypoint.sh"]

CMD ["up"]