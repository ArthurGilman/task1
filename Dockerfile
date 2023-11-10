FROM golang:latest AS builder

WORKDIR /src

COPY ./ ./

RUN GOOS=linux go build -o /main main.go

FROM ubuntu:23.04

WORKDIR /app

COPY --from=builder main main

COPY --from=builder src/db.json app/db.json

COPY --from=builder src/db.csv app/db.csv

CMD ["./main", "app/db.csv"]




