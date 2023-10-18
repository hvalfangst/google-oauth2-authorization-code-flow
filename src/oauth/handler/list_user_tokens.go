package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/repository"
)

func ListUserTokens(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userTokens, err := repository.ListUserTokens(db)
		if err != nil {
			c.JSON(404, gin.H{"error": "Could not list UserTokens"})
			return
		}

		if len(userTokens) == 0 {
			c.JSON(200, gin.H{"message": "No UserTokens were found"})
			return
		}

		c.JSON(200, gin.H{"user_tokens": userTokens})
	}
}
