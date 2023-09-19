run:
	go run ./cmd/api

swag:
	swag init -g ./cmd/api/main.go

start:
	sudo docker start 168440bcc15d

stop:
	sudo docker stop 168440bcc15d

status:
	sudo docker container ls -a