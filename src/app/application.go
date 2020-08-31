package app

import (
	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_oauth-api/src/clients/cassandra"
	"github.com/migueloli/bookstore_oauth-api/src/domais/accesstoken"
	"github.com/migueloli/bookstore_oauth-api/src/http"
	"github.com/migueloli/bookstore_oauth-api/src/repository/db"
)

var router = gin.Default()

// StartApplication function to prepare and configure the application
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
