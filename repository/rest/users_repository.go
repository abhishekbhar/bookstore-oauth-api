package rest

import (
	"time"
	"encoding/json"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
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
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

func NewRepository() RestUsersRepository{
	return &usersRepository{}
}

type usersRepository struct {}

func (u *usersRepository)LoginUser(email, password string) (*users.User, *rest_errors.RestErr) {
	request  := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("Invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal user")
	}

	return &user, nil
}
