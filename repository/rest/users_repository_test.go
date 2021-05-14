package rest

import (
	"os"
	"fmt"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/golang-restclient/rest"
) 

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases.")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: 			"localhost:8080/users/login",
		HTTPMethod: 	http.MethodPost,
		ReqBody: 		`{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: 	-1,
		RespBody:		`{}`,
	})
	repository := usersRepository{}
	user, err  := repository.LoginUser("email@gmail.com","the-password") 
	assert.Nil(t, user)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: 			"localhost:8080/users/login",
		HTTPMethod: 	http.MethodPost,
		ReqBody: 		`{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: 	http.StatusNotFound,
		RespBody:		`{"message":"invalid login credentials", "status":"404","error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err  := repository.LoginUser("email@gmail.com","the-password") 
	assert.Nil(t, user)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest client response when trying to login user", err.Message)
	
}


func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: 			"localhost:8080/users/login",
		HTTPMethod: 	http.MethodPost,
		ReqBody: 		`{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: 	http.StatusNotFound,
		RespBody:		`{"message":"invalid login credentials", "status":404,"error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err  := repository.LoginUser("email@gmail.com","the-password") 
	assert.Nil(t, user)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: 			"localhost:8080/users/login",
		HTTPMethod: 	http.MethodPost,
		ReqBody: 		`{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: 	http.StatusOK,
		RespBody:		`{"id": "1", "first_name":"Abhsihek", "last_name":"Bhardwaj","email":"volage.abhishek@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err  := repository.LoginUser("email@gmail.com","the-password") 
	assert.Nil(t, user)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal user", err.Message)
	
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: 			"localhost:8080/users/login",
		HTTPMethod: 	http.MethodPost,
		ReqBody: 		`{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: 	http.StatusOK,
		RespBody:		`{"id": 1, "first_name":"Abhsihek", "last_name":"Bhardwaj","email":"volage.abhishek@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err  := repository.LoginUser("email@gmail.com","the-password") 
	assert.Nil(t, err)
	assert.NotNil(t,user)
	// assert.EqualValues(t, http.StatusOK, err.Status)
}