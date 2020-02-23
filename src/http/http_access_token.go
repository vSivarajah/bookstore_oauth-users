package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vSivarajah/bookstore_oauth-users/src/domain/access_token"
	"github.com/vSivarajah/bookstore_oauth-users/src/services/access_tokens"
	"github.com/vSivarajah/bookstore_users-api/utils/errors"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_tokens.Service
}

func NewHandler(service access_tokens.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var accessTokenRequest access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&accessTokenRequest); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	accessToken, err := handler.service.Create(accessTokenRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
