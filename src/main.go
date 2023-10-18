package main

import (
	"github.com/gin-gonic/gin"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/common/configuration"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/common/db"
	Oauth2 "hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/route"
	"log"
)

func main() {
	r := gin.Default()

	// Fetch JSON based on key 'db' for file 'configuration.json'
	conf, err := configuration.Get("db")
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'configuration.json'
	database := db.ConnectDatabase(conf.(configuration.Db))
	defer db.CloseDatabase(database)

	// Creates the table: 'user_tokens'
	db.CreateUserTokensTable(err, database)

	// Configures route associated with Oauth2 Google IDP
	Oauth2.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
