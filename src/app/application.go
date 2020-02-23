package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vSivarajah/bookstore_oauth-users/src/clients/cassandra"
	"github.com/vSivarajah/bookstore_oauth-users/src/http"
	"github.com/vSivarajah/bookstore_oauth-users/src/repository/db"
	"github.com/vSivarajah/bookstore_oauth-users/src/repository/rest"
	"github.com/vSivarajah/bookstore_oauth-users/src/services/access_tokens"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	defer session.Close()

	accessTokenHandler := http.NewHandler(access_tokens.NewService(rest.NewRestUsersRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetById)
	router.POST("/oauth/access_token", accessTokenHandler.Create)
	router.Run(":8080")
}
