package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strconv"
	"wallet-api/database"
	"wallet-api/models"
)

type Body struct {
	Amount float64 `json:"amount" binding:"required"`
}

//Ping - pings the server and returns a response
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetWalletBalance - returns the balance of the requested wallet id
func GetWalletBalance(c *gin.Context) {
	wallet := models.Wallet{}

	// get the wallet_id as route params
	walletId := c.Param("wallet_id")
	if walletId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wallet ID was not supplied",
		})
		return
	}
	key := "walletID_" + walletId
	balance, err := database.RDB.Get(database.CTX, key).Result()

	if balance != "" && err == nil {
		c.JSON(200, gin.H{
			"message": "Successfully returned balance for wallet",
			"balance": balance,
		})
		return
	}

	// get the wallet information using the wallet id
	err = database.DB.First(&wallet, walletId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Could not find the Wallet with the specified wallet_id",
		})
		return
	}

	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "System encountered an error"})
		return
	}

	// return response with wallet balance
	c.JSON(200, gin.H{
		"message": "Successfully returned balance for wallet",
		"balance": wallet.Balance,
	})
	return
}

// CreditWallet - credits the wallet balance with the requested amount
func CreditWallet(c *gin.Context) {
	wallet := models.Wallet{}

	// get the wallet_id as route params
	walletId := c.Param("wallet_id")
	if walletId == "" {
		c.JSON(400, gin.H{
			"message": "Wallet ID was not supplied",
		})
		return
	}
	body := Body{}

	// extract the body from the request body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Amount is required",
		})
		return
	}
	// convert the amount the decimal
	amount := decimal.NewFromFloat(body.Amount)

	// check if the amount to be credited is negative
	if amount.IsNegative() {
		c.JSON(404, gin.H{
			"message": "Amount to be credited cannot be negative",
		})
		return
	}
	// check if the amount to be credit to be credited is zero
	if amount.Equals(decimal.Zero) {
		c.JSON(404, gin.H{
			"message": "Amount to be credited cannot be zero",
		})
		return
	}

	// get the wallet information using the wallet id
	err = database.DB.First(&wallet, walletId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Could not find the Wallet with the specified wallet_id",
		})
		return
	}

	//add the amount to be credited to the wallet balance
	wallet.Balance, _ = decimal.NewFromFloat(wallet.Balance).Add(amount).Float64()

	// update the DB
	database.DB.Save(&wallet)

	// update the redis cache with new balance
	key := "walletID_" + strconv.FormatUint(uint64(wallet.ID), 10)
	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "System encountered an error",
		})
		return
	}

	// return the success response and updated balance
	c.JSON(200, gin.H{
		"message": "Successfully credited wallet",
		"balance": wallet.Balance,
	})
	return
}

// DebitWallet - debits the requested amount from the wallet balance
func DebitWallet(c *gin.Context) {
	wallet := models.Wallet{}

	// get the wallet_id as route params
	walletId := c.Param("wallet_id")
	if walletId == "" {
		c.JSON(400, gin.H{
			"message": "Wallet ID was not supplied",
		})
	}
	body := Body{}

	// extract the amount from the request body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Amount is required",
		})
		return
	}
	// convert the amount the Decimal
	amount := decimal.NewFromFloat(body.Amount)

	// check if the amount to be debited is zero
	if amount.Equals(decimal.Zero) {
		c.JSON(404, gin.H{
			"message": "Amount to be debited cannot be zero",
		})
		return
	}

	// check if the amount to be debited is negative
	if amount.IsNegative() {
		c.JSON(404, gin.H{
			"message": "Amount to be debited cannot be negative",
		})
		return
	}

	// get the wallet information using the wallet id
	err = database.DB.First(&wallet, walletId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Could not find the Wallet with the specified wallet_id",
		})
		return
	}

	// subtract the amount to be credited to the wallet balance
	isBalanceNegative := decimal.NewFromFloat(wallet.Balance).Sub(amount).IsNegative()

	if isBalanceNegative {
		c.JSON(400, gin.H{
			"message": "Wallet balance is not sufficient to perform transaction",
		})
		return
	}
	wallet.Balance, _ = decimal.NewFromFloat(wallet.Balance).Sub(amount).Float64()

	// update the DB
	database.DB.Save(&wallet)

	// update the redis cache with new balance
	key := "walletID_" + strconv.FormatUint(uint64(wallet.ID), 10)

	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "System encountered an error",
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "Successfully debited wallet",
		"balance": wallet.Balance,
	})
}
