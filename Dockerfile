
FROM docker.io/library/golang:1.21.0

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/api

CMD ["app"] 