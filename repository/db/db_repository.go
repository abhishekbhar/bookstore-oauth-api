package db

import (
	"github.com/abhishekbhar/bookstore-oauth-api/utils/errors"
	"github.com/abhishekbhar/bookstore-oauth-api/domain/access_token"
	"github.com/abhishekbhar/bookstore-oauth-api/clients/cassandra"
)




func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {}

const (
	queryCreateAccessToken    = "INSERT INTO access_token (access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryGetAccessToken 	  = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
	queryUpdateExpirationTime = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr){

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires) ; err != nil {
			if err.Error() == "not found" {
				return nil, errors.NewBadRequestError("no access token found with given id")
			}
			return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}


func (dbr *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {


	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec() ; err != nil {
			return errors.NewInternalServerError(err.Error())
	}
	return nil
}


func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpirationTime,
		at.Expires,
		at.AccessToken,
		).Exec() ; err != nil {
			return errors.NewInternalServerError(err.Error())
	}

	return nil
}
