package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wallet-api/models"
	"wallet-api/utils"
)

var DB *gorm.DB

// ConnectDB init connection to MySQL database
func ConnectDB(databaseURL string) {
	var err error
	DB, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the Models
	err = DB.AutoMigrate(&models.Player{}, &models.Wallet{})
	if err != nil {
		panic(err)
	}

}

func SeedPlayers() {
	users := []models.Player{{
		FirstName:    "Tony",
		LastName:     "Cookey",
		EmailAddress: "tony@example.com",
		Password:     utils.HashPassword("password123"),
	}, {
		FirstName:    "Chris",
		LastName:     "Ronney",
		EmailAddress: "chris@example.com",
		Password:     utils.HashPassword("password123"),
	}}

	result := DB.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	}
}
func SeedWallets() {
	wallets := []models.Wallet{{
		PlayerID: 1,
		Balance:  1000,
	}, {
		PlayerID: 2,
		Balance:  900,
	}}

	result := DB.Create(&wallets)

	if result.Error != nil {
		panic(result.Error)
	}
}
