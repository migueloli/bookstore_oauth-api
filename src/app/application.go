package app

import (
	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
	"github.com/migueloli/bookstore_oauth-api/src/http"
	"github.com/migueloli/bookstore_oauth-api/src/repository/db"
)

var router = gin.Default()

// StartApplication function to prepare and configure the application
func StartApplication() {
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
