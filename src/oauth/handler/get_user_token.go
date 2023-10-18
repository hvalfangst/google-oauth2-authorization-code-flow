package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/repository"
	"strconv"
)

func GetUserToken(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string value associated with key 'id' and convert it to Integer
		IDStringParam := c.Param("id")
		ID, err := strconv.ParseInt(IDStringParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UserToken ID"})
			return
		}

		// Query table 'user_tokens' by ID
		token, err := repository.GetUserToken(db, ID)
		if err != nil {
			c.JSON(404, gin.H{"error": "UserToken doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"user_token": token})
	}
}
