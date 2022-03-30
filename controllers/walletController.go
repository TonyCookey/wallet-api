package controllers

import (
	"github.com/gin-gonic/gin"
	"wallet-api/database"
	"wallet-api/models"
)

type Response struct {
	message string
	result  struct{}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetWalletBalance(c *gin.Context) {
	wallet := models.Wallet{}

	// get the wallet_id as route params
	walletId := c.Param("wallet_id")
	if walletId == "" {
		c.JSON(400, gin.H{
			"message": "Wallet ID was not supplied",
		})
	}
	key := "walletID_" + walletId
	balance, err := database.RDB.Get(database.CTX, key).Result()

	if balance != "" {
		c.JSON(200, gin.H{
			"message": "Successfully returned balance for wallet from redis",
			"balance": balance,
		})
		return
	}

	// get the wallet information using the wallet id
	database.DB.First(&wallet, walletId)

	if wallet.UserID == 0 || wallet.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Could not find the Wallet with the specified wallet_id",
		})
		return
	}

	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		panic(err)
	}

	// return response with wallet balance
	c.JSON(200, gin.H{
		"message": "Successfully returned balance for wallet",
		"balance": wallet.Balance,
	})
	return
}
