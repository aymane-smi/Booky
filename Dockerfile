FROM postgres:latest

COPY book_go.sql .

ENV POSTGRES_USER="aymane"
ENV POSTGRES_PASSWORD="aymane@123"
ENV POSTGRES_DB="book_go"

EXPOSE 5432