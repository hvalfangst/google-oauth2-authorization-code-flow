# OAuth2-authorization-flow-using-Google

## Abstract


## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)

## Startup

The script "up" provisions resources and starts our application by executing the following:
```
1. docker-compose -f docker/db/docker-compose.yml up -d
2. go build -o gin_api_with_auth src/main.go
3. ./gin_api_with_auth
```

## Shutdown

The script "down" deletes our dev and test databases by executing the following:
```
1. docker-compose -f db/docker-compose.yml down
```

## Postman Collection

The repository includes a Postman collection in the 'postman' directory.