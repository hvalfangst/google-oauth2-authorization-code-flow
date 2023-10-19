# OAuth2-authorization-flow-using-Google

## Abstract
This project provides a basic example in which OAuth2 authorization code flow is being utilized in order to access Google resources on behalf of the user. 
User and token information is persisted to DB upon successful authorization.


## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)
* [Google account](https://accounts.google.com/signup/v2/createaccount?theme=glif&flowName=GlifWebSignIn&flowEntry=SignUp)


## OAuth2 Credentials

For this project to work, you need to replace the 'PLACEHOLDER' values associated with keys "ClientID" and "ClientSecret" in file 'configuration.json'.

In order to get hold of Client ID and Secret, one must do the following:

1. Access the [credentials section](https://console.cloud.google.com/apis/credentials) under Google APIs & Services. Create a new project if you haven't done so already.
2. Click on the button with label 'CREATE CREDENTIALS'
3. Choose option 'OAuth client ID'
4. Choose option 'Web application'
5. Add the following under 'Authorized redirect URIs': http://localhost:8080/auth/google/callback
6. Click on the 'CREATE' button

Now you can replace fields contained in file 'configuration.json' with your actual ID & Secret

## Application Flow

1. User visits the (static HTML) login page by accessing the following URL in the browser: http://localhost:8080
2. User clicks on the "Login with Google" button, which redirects to http://localhost:8080/auth/google/login 
3. The application initiates the OAuth2 Authorization Code Flow, redirecting users to Google's authentication page.
4. Users log in to their Google accounts and authorize the application to access their Google account information.
5. Upon successful authorization, Google sends an authorization code to the application's callback URL at http://localhost:8080/auth/google/callback
6. The application exchanges the authorization code contained in callback response for an access token and other necessary credentials.
7. The access token is used to make requests to Google's UserInfo API, which in this case retrieves email associated with user.
8. The application persists a record containing the user's access token in DB.


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