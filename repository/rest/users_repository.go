package rest

import (
	"time"
	"encoding/json"
	"github.com/abhishekbhar/bookstore-oauth-api/utils/errors"
	"github.com/abhishekbhar/bookstore-oauth-api/domain/users"
	"github.com/mercadolibre/golang-restclient/rest"

)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

func NewRepository() RestUsersRepository{
	return &usersRepository{}
}

type usersRepository struct {}

func (u *usersRepository)LoginUser(email, password string) (*users.User, *errors.RestErr) {
	request  := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user")
	}

	return &user, nil
}
