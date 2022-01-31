FROM postgres:latest

COPY ./pkg/data/db/init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 5432
