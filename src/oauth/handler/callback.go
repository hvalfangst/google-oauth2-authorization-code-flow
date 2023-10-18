package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/google"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/model"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/repository"
	"log"
	"net/http"
	"time"
)

func Callback(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Read oauthState from Cookie
		oauthState, err := c.Request.Cookie("oauthstate")
		if err != nil {
			log.Println(err)
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		// Ensure that the state objects are identical in order to detect CSRF attacks
		if c.DefaultQuery("state", "nil") != oauthState.Value {
			log.Println("Invalid oauth google state detected: Potential CSRF")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		// Utilize code derived from response to make a token request, which then in turn is used to fetch user data
		rawUserData, token, err := google.GetUserData(c.DefaultQuery("code", ""))
		if err != nil {
			log.Println(err)
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		// Unmarshal raw bytes into JSON
		jsonData, err, jsonString, done := bytesToJSON(err, rawUserData)
		if done {
			return
		}

		fmt.Println(string(jsonString))

		// Populate the UserToken struct with values derived from JSON
		userToken := model.UserToken{
			UserID:       jsonData["id"].(string),
			UserEmail:    jsonData["email"].(string),
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			TokenExpiry:  token.Expiry,
			TokenType:    token.TokenType,
			Provider:     "Google",
			CreatedAt:    time.Now(),
		}

		err = repository.CreateUserToken(db, &userToken)

		if err != nil {
			log.Printf("Creation of UserToken failed with the following error: %v\n\n\n", err)
			return
		}

		c.String(http.StatusOK, "UserInfo: %s\n", rawUserData)
	}
}

func bytesToJSON(err error, rawData []byte) (map[string]interface{}, error, []byte, bool) {

	// Create a map to unmarshal the JSON data
	var jsonData map[string]interface{}

	// Unmarshal the byte data into the map
	err = json.Unmarshal(rawData, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil, nil, true
	}

	// Marshal the map back to a JSON string
	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil, nil, true
	}
	return jsonData, err, jsonString, false
}
