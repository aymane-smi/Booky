FROM postgres:latest

WORKDIR /db

COPY ./book_go.sql .

ENV POSTGRES_USER="aymane"
ENV POSTGRES_PASSWORD="aymane@123"
ENV POSTGRES_DB="book_go"

RUN psql -h localhost -U aymane book_go -a -f book_go.sql

EXPOSE 5432