package http

import (
	// "strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/abhishekbhar/bookstore-oauth-api/domain/access_token"
	"github.com/abhishekbhar/bookstore-oauth-api/utils/errors"	
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}


type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service: service,
	}	
}


func (h *accessTokenHandler)GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}


func (h *accessTokenHandler) Create(c *gin.Context) {
	var at *access_token.AccessToken
	var restErr *errors.RestErr
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr = errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	if at, restErr= h.service.Create(atr); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, at)
}
