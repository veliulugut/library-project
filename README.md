# Library
Library is a book library backend application that manage books.

## Usage 
You can checkout makefile file for run commands.

### Running with docker 
```shell
docker build . -t library
```
```shell
docker run --name libraryexample -p 8080:8080 -e SMTP_HOST=<host> -e SMTP_PORT=<port> -e SMTP_FROM=<email> -e SMTP_PASSWORD=<password> library
```

## Used Libraries

- Gin Web Framework (v1.9.1)
- Swaggo swagger generator (v1.16.2)
- Redis (v9.1.0)
- Entgo(ORM) (v0.12.3)
- Golang JWT (v5.0.0)
---
Swagger api endpoint: `http://localhost:8080/swagger/index.html`
