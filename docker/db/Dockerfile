# Dockerfile
FROM postgres:latest

ENV POSTGRES_DB=nss
ENV POSTGRES_USER=lafin
ENV POSTGRES_PASSWORD=lafin1234

COPY init.sql /docker-entrypoint-initdb.d/