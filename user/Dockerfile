FROM alpine:3.18.4

WORKDIR /app

COPY ./bin/user /app/user
COPY ./config /app/config

EXPOSE 8081

ENTRYPOINT ["./user"]
