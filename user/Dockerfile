FROM alpine:3.18.4

WORKDIR /app

COPY ./bin/user /app/user
COPY ./bin/migrate /app/migrate
COPY ./config /app/config

EXPOSE 8081

ENTRYPOINT ["./user"]
