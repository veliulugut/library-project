run:
	go run ./cmd/api

run-test-data:
	TEST_DATA=true go run ./cmd/api

swag:
	swag init -g ./cmd/api/main.go

start:
	sudo docker start 168440bcc15d

stop:
	sudo docker stop 168440bcc15d

status:
	sudo docker container ls -a

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature "sql/upsert","sql/execquery" ./ent/schema