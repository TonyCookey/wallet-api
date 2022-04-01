package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"wallet-api/controllers"
	"wallet-api/middleware"
)

func InitializeRoutes() *gin.Engine {
	server := gin.New()

	// use sirupsen logger with gin
	log := logrus.New()
	server.Use(ginlogrus.Logger(log), gin.Recovery())

	// define api v1 routes
	api := server.Group("api/v1")
	{
		api.GET("/ping", controllers.Ping)
		api.GET("/seed", controllers.SeedDB)
		api.POST("/login", controllers.Login)

		//group wallet routes which are protected and require authentication
		wallet := api.Group("wallets")

		// use AuthorizeJWT middleware
		wallet.Use(middleware.AuthorizeJWT())

		wallet.GET("/:wallet_id/balance", controllers.GetWalletBalance)
		wallet.POST("/:wallet_id/credit", controllers.CreditWallet)
		wallet.POST("/:wallet_id/debit", controllers.DebitWallet)
	}
	server.NoRoute(noRoute)

	return server
}
func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Route not found"})
}
