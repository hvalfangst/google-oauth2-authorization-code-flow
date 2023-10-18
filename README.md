# OAuth2-authorization-flow-using-Google

## Abstract
This project implements an application that utilizes OAuth2 Authorization Code Flow with Google as the identity provider. The application provides a user-friendly way for users to log in with their Google accounts. Upon login, the application performs a callback to the back-channel, where the code associated with the callback response is exchanged for an access token. The token information is persisted to a database upon successful exchange. Additionally, the application includes dedicated endpoints for retrieving and deleting tokens.


## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)


## OAuth2 Credentials

For this project to work, you need to replace the 'PLACEHOLDER' values associated with keys "ClientID" and "ClientSecret" in file 'configuration.json'.

In order to get hold of Client ID and Secret, one must do the following:

1. Access the [credentials section](https://console.cloud.google.com/apis/credentials) under Google APIs & Services. Create a new project if you haven't done so already.
2. Click on the button with label 'CREATE CREDENTIALS'
3. Choose option 'OAuth client ID'
4. Choose option 'Web application'
5. Under 'Authorized redirect URIs', add the following: 'localhost:8080/auth/google/callback'
6. Click on the 'CREATE' button

Now you can replace fields contained in file 'configuration.json' with your newly generated client ID and secret by act of copy pasta :)


## Startup

The script "up" provisions resources and starts our application by executing the following:
```
1. docker-compose -f docker/db/docker-compose.yml up -d
2. go build -o OAuth2-authorization-code-flow-using-Google src/main.go
3. ./OAuth2-authorization-code-flow-using-Google
```

## Shutdown

The script "down" deletes our dev and test databases by executing the following:
```
1. docker-compose -f db/docker-compose.yml down
```

## Postman Collection

The repository includes a Postman collection in the 'postman' directory.