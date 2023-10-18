package repository

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/OAuth2-authorization-code-flow-using-Google/src/oauth/model"
	"log"
)

func CreateUserToken(db *pg.DB, token *model.UserToken) error {
	_, err := db.Model(token).Insert()
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return err
	}
	return nil
}

func GetUserToken(db *pg.DB, ID int64) (*model.UserToken, error) {
	token := &model.UserToken{}
	err := db.Model(token).Where("id = ?", ID).Select()
	if err != nil {
		log.Printf("Error retrieving token by ID: %v", err)
		return nil, err
	}
	return token, nil
}

func ListUserTokens(db *pg.DB) ([]*model.UserToken, error) {
	var tokens []*model.UserToken
	err := db.Model(&tokens).Select()
	if err != nil {
		log.Printf("Error retrieving tokens: %v", err)
		return nil, err
	}
	return tokens, nil
}

func DeleteUserToken(db *pg.DB, ID int64) error {
	token := &model.UserToken{ID: ID}
	_, err := db.Model(token).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting token: %v", err)
		return err
	}
	return nil
}
