FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:bullseye-slim

COPY --from=build /app/main /main

ENTRYPOINT ["/main"]

EXPOSE 3000
