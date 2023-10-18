package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {

	// Route associated with static HTML login page
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/auth/google/login", handler.Login())
	r.GET("/auth/google/callback", handler.Callback(db))
	r.GET("/user-tokens", handler.ListUserTokens(db))
	r.GET("/user-tokens/:id", handler.GetUserToken(db))
	r.DELETE("/user-tokens/:id", handler.DeleteUserToken(db))
}
