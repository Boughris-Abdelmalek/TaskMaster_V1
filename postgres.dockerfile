FROM postgres:latest

LABEL author="Malik"
LABEL description="Postgres image for task master"
LABEL version="1.0"

COPY migrations/*.sql /docker-entrypoint-initdb.d/