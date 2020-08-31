package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
)

// AccessTokenHandler is a struct for the handler.
type AccessTokenHandler interface {
	GetByID(*gin.Context)
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
