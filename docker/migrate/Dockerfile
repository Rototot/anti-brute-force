FROM migrate/migrate:latest

WORKDIR /app/migrations

COPY ./docker/migrate/entrypoint.sh /opt/
COPY ./db/migrations /app/migrations

ENV APP_POSTGRES_HOST=app
ENV APP_POSTGRES_DB=app
ENV APP_POSTGRES_USER=app
ENV APP_POSTGRES_PASSWORD=app_pass
ENV APP_POSTGRES_PORT=5432

ENTRYPOINT ["/opt/entrypoint.sh"]

CMD ["up"]