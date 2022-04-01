package controllers

import (
	"github.com/gin-gonic/gin"
	"wallet-api/models"
	"wallet-api/services"
)

var jwtService services.JWTService = services.JWTAuthService()

func Login(c *gin.Context) {
	player := models.Player{}
	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "",
			"player":  player,
		})
		return
	}

	isUserAuthenticated := services.ValidateUser(player.EmailAddress, player.Password)
	if !isUserAuthenticated {
		c.JSON(404, gin.H{
			"message": "",
		})
		return
	}
	token := jwtService.GenerateToken(player.EmailAddress, true)

	c.JSON(200, gin.H{
		"token": token,
	})
	return
}
