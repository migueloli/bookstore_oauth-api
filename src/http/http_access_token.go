package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
	"github.com/migueloli/bookstore_oauth-api/src/utils/errors"
)

// AccessTokenHandler is a struct for the handler.
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewHandler cria um handler http.
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	at := accesstoken.AccessToken{}
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
