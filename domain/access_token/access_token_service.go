package access_token

import (
	"strings"
	"github.com/abhishekbhar/bookstore-oauth-api/domain/users"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
)

type DBRepository interface{
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}


type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}

type service struct {	
	dbRepo DBRepository
	restUsersRepo RestUsersRepository
}


func NewService(repo DBRepository, userRepo RestUsersRepository) Service {
	return &service{
		dbRepo: repo,
		restUsersRepo: userRepo,
	}
}


func (s *service) GetById(id string) (*AccessToken, *rest_errors.RestErr) {
	accessToken := strings.TrimSpace(id)
	if len(accessToken) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	return s.dbRepo.GetById(id)
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *rest_errors.RestErr) {

	if err:= request.Validate(); err != nil {
		return nil,err
	}

	// TODO: support both client credentials and password grant type
	user, err := s.restUsersRepo.LoginUser(request.UserName, request.Password)
	if err!= nil {
		return nil,err
	}

	at := GetNewAccessToken(user.Id)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at,nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *rest_errors.RestErr {
	if err:= at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}

