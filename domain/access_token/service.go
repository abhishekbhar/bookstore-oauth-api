package access_token

import (
	"github.com/abhishekbhar/bookstore-oauth-api/utils/errors"
)

type Repository interface{
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)

}

type service struct {	
	repository Repository
}


func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}


func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	return s.repository.GetById(id)
}