package access_token

import (
	"time"
	"fmt"
	"strings"
	"github.com/abhishekbhar/bookstore-utils-go/rest_errors"
	"github.com/abhishekbhar/bookstore-oauth-api/utils/crypto_utils"
)

const (
	expirationTime 				= 24
	grantTypePassword 			= "password"
	grantTypeClientCredentials 	= "client_credentials"
)


type AccessTokenRequest struct {
	GrantType	string	`json:"grant_type"`
	Scope		string	`json:"scope"`

	// Used for password grant types
	UserName	string	`json:"username"`
	Password	string	`json:"password"`

	// Used for client credentials grant type
	ClientId	string	`json:"client_id"`
	ClientSecret string `json:"client_secret"`
}


func (atr *AccessTokenRequest) Validate() *rest_errors.RestErr {

	switch atr.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

	//Validate parameters for each grant_type
	return nil
}




type AccessToken struct {
	AccessToken string  `json:"access_token"`
	UserId 		int64   `json:"user_id"`
	ClientId	int64   `json:"client_id"`	
	Expires		int64	`json:"expires"`	
}


func GetNewAccessToken(id int64) AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),

	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0, )
	return expirationTime.Before(now)
}


func (at *AccessToken) Validate() *rest_errors.RestErr{
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil


}



func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId,at.Expires))
}
