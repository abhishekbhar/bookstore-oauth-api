package db

import (
	"github.com/abhishekbhar/bookstore-oauth-api/domain/access_token"
	"github.com/abhishekbhar/bookstore-oauth-api/utils/errors"
)


func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {}



func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr){
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}
