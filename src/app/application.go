package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vSivarajah/bookstore_oauth-users/src/clients/cassandra"
	"github.com/vSivarajah/bookstore_oauth-users/src/domain/access_token"
	"github.com/vSivarajah/bookstore_oauth-users/src/http"
	"github.com/vSivarajah/bookstore_oauth-users/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	accessTokenHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetById)

	router.Run(":8080")
}
