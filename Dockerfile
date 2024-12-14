FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env.docker .env

RUN go build -o api cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api ./
COPY .env.docker .env
COPY deploy deploy

EXPOSE 8080
CMD ["./api"]
