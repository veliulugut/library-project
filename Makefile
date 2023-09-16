run:
	go run ./cmd/api

swag:
	swag init -g ./cmd/api/main.go

