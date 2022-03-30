package routes

import (
	"github.com/gin-gonic/gin"
	"wallet-api/controllers"
)

func InitializeRoutes() {
	router := gin.Default()

	// define api v1 routes
	v1 := router.Group("api/v1")
	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/wallets/:wallet_id/balance", controllers.GetWalletBalance)

	}
	router.NoRoute(noRoute)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := router.Run()
	if err != nil {
		return
	}
}
func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Route not found"})
}
