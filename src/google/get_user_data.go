package google

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/model"
	"io"
	"log"
	"net/http"
)

func GetUserData(code string) ([]byte, *oauth2.Token, error) {

	// Use code to get a token from Google.
	token, err := model.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	// Check if the token is valid.
	if token.Valid() {

		// Get user info from Google using the valid access token.
		response, err := http.Get(model.OauthGoogleUrlAPI + token.AccessToken)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get user info from Google: %v", err)
		}

		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		return bytes, token, nil
	}

	return nil, nil, fmt.Errorf("invalid token received from Google")
}
