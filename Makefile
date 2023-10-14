run:
	go run ./cmd/api

run-test-data:
	TEST_DATA=true go run ./cmd/api

swag:
	swag init -g ./cmd/api/main.go

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature "sql/upsert","sql/execquery" ./ent/schema

