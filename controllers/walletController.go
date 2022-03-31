package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"wallet-api/database"
	"wallet-api/models"
)

type Body struct {
	Amount string `json:"amount"`
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
	}
	key := "walletID_" + walletId
	balance, err := database.RDB.Get(database.CTX, key).Result()

	if balance != "" {
		c.JSON(200, gin.H{
			"message": "Successfully returned balance for wallet",
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
		c.JSON(500, gin.H{
			"message": "System encountered an error",
			"error":   err,
		})
		panic(err)
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
	}
	body := Body{}
	// extract the body from the request body
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	// convert the amount the decimal
	amount, err := decimal.NewFromString(body.Amount)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Amount is required",
			"error":   err,
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
	database.DB.First(&wallet, walletId)

	//add the amount to be credited to the wallet balance
	wallet.Balance, _ = decimal.NewFromFloat(wallet.Balance).Add(amount).Float64()

	// update the DB
	database.DB.Save(&wallet)

	key := "walletID_" + strconv.FormatUint(uint64(wallet.ID), 10)

	// update the redis cache with new balance
	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(400, gin.H{
		"message": "Successfully credited wallet",
		"balance": wallet.Balance,
	})
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

	// extract the body from the request body
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	// convert the amount the decimal
	amount, err := decimal.NewFromString(body.Amount)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Amount is required",
			"error":   err,
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
	database.DB.First(&wallet, walletId)

	//add the amount to be credited to the wallet balance
	wallet.Balance, _ = decimal.NewFromFloat(wallet.Balance).Sub(amount).Float64()

	// update the DB
	database.DB.Save(&wallet)

	// update the redis cache with new balance
	key := "walletID_" + strconv.FormatUint(uint64(wallet.ID), 10)

	err = database.RDB.Set(database.CTX, key, wallet.Balance, 0).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(400, gin.H{
		"message": "Successfully debit wallet",
		"balance": wallet.Balance,
	})
}
