
FROM docker.io/library/golang:1.21.0 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -v -o /app/api ./cmd/api


FROM docker.io/library/alpine:3.14.2

WORKDIR /app

COPY --from=builder /app/api /app/api


CMD ["/app/api"]