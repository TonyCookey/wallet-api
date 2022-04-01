package controllers

import (
	"github.com/gin-gonic/gin"
	"wallet-api/database"
)

// SeedDB - Seeds data into the database
func SeedDB(c *gin.Context) {

	database.SeedPlayers()
	database.SeedWallets()

	c.JSON(200, gin.H{
		"message": "seeded players and wallets into DB",
	})
}
