package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"wallet-api/database"
)

func TestSetupServer(t *testing.T) {
	err := godotenv.Load("../.env.testing")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB(os.Getenv("DATABASE_URL"))

	database.InitializeRedisInstance(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"), 0)
}

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Ping(c)

	assert.Equal(t, 200, w.Code)
}
func TestDebitWalletNoWalletID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	DebitWallet(c)

	assert.Equal(t, 400, w.Code)
}

func TestCreditWalletNoWalletID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	CreditWallet(c)

	assert.Equal(t, 400, w.Code)

}

func TestGetWalletBalanceWalletNoWalletID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   "wallet_id",
			Value: "1",
		},
	}
	GetWalletBalance(c)

	assert.Equal(t, 200, w.Code)

}
