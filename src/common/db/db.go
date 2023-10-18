package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/common/configuration"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/model"
	"log"
)

func ConnectDatabase(config configuration.Db) *pg.DB {
	opts := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Address,
		Database: config.Database,
	}
	return pg.Connect(opts)
}

func CreateTable(db *pg.DB, model interface{}) error {
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateUserTokensTable Creates the 'user_tokens' table
func CreateUserTokensTable(err error, database *pg.DB) {
	err = CreateTable(database, (*model.UserToken)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
}

func CloseDatabase(db *pg.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
