package routes

import (
	"github.com/gin-gonic/gin"
	"wallet-api/controllers"
	"wallet-api/middleware"
)

func InitializeRoutes() {
	//var loginService services.LoginService = services.StaticLoginService()
	//var jwtService services.JWTService = services.JWTAuthService()
	//var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	router := gin.Default()

	// define api v1 routes
	api := router.Group("api/v1")
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
	router.NoRoute(noRoute)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Route not found"})
}
