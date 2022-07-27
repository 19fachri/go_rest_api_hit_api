package main

import (
	"testhex10/rest_api/controllers"
	"testhex10/rest_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/public/:wallet_address/contents", controllers.GetTokenByWallet)
	router.POST("/public/fetch_tokens", controllers.PostTokenByWallet)
	router.DELETE("/public/token/delete", controllers.DeleteToken)

	router.Run("localhost:8080")
}
