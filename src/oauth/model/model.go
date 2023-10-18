package model

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/common/configuration"
	"time"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     getClientID(),
	ClientSecret: getClientSecret(),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func getClientID() string {
	googleConfig, err := configuration.Get("google")
	if err != nil {
		panic("Error: Unable to retrieve Google configuration")
	}

	clientID := googleConfig.(configuration.Google).ClientID
	if clientID == "PLACEHOLDER" {
		panic("Error: Missing ClientID in 'configuration.json'")
	}

	return clientID
}

func getClientSecret() string {
	googleConfig, err := configuration.Get("google")
	if err != nil {
		panic("Error: Unable to retrieve Google configuration")
	}

	clientID := googleConfig.(configuration.Google).ClientSecret
	if clientID == "PLACEHOLDER" {
		panic("Error: Missing ClientID in 'configuration.json'")
	}

	return clientID
}

type UserToken struct {
	ID           int64     `json:"id"`
	UserID       string    `json:"user_id"`
	UserEmail    string    `json:"user_email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenExpiry  time.Time `json:"token_expiry"`
	TokenType    string    `json:"token_type"`
	Provider     string    `json:"provider"`
	CreatedAt    time.Time `json:"created_at"`
	_            struct{}  `pg:"_schema:user_tokens"`
}
