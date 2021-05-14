package app

import (
	"github.com/abhishekbhar/bookstore-oauth-api/domain/access_token"
	"github.com/abhishekbhar/bookstore-oauth-api/repository/db"
	"github.com/abhishekbhar/bookstore-oauth-api/repository/rest"
	"github.com/abhishekbhar/bookstore-oauth-api/http"
	"github.com/gin-gonic/gin"
)


var (
	router = gin.Default()
)

func StartApp() {
	atService := access_token.NewService(db.NewRepository() , rest.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}