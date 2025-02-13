FROM golang:1.21.0 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /app/api ./cmd/api

FROM alpine:3.14.2

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/api /usr/local/bin/api

ENTRYPOINT [ "api" ]