package handler

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/model"
	"net/http"
	"time"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		oauthState, err := GenerateOAuthStateCookie(c.Writer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate OAuth state cookie",
			})
			return
		}

		u := model.GoogleOauthConfig.AuthCodeURL(oauthState)
		c.Redirect(http.StatusFound, u)
	}
}

// GenerateOAuthStateCookie generates an OAuth state cookie and sets it in the HTTP response.
func GenerateOAuthStateCookie(w http.ResponseWriter) (string, error) {

	// Generate a random 16-byte slice for the state
	stateBytes := make([]byte, 16)
	_, err := rand.Read(stateBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a base64 URL-encoded string
	state := base64.URLEncoding.EncodeToString(stateBytes)

	// Set the cookie with an expiration time
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  expiration,
		HttpOnly: true, // Enforce HttpOnly for added security
	}

	http.SetCookie(w, &cookie)

	return state, nil
}
