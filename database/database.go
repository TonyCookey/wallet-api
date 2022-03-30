package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wallet-api/models"
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
	err = DB.AutoMigrate(&models.User{}, &models.Wallet{})
	if err != nil {
		panic(err)
	}

	// Seed data into the database
	seedUsers()
	seedWallets()
}

func seedUsers() {
	users := []models.User{{
		FirstName:    "Tony",
		LastName:     "Cookey",
		EmailAddress: "tony@example.com",
	}, {
		FirstName:    "Tonero",
		LastName:     "Cookey",
		EmailAddress: "tonero@example.com",
	}}

	result := DB.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	}
}
func seedWallets() {
	wallets := []models.Wallet{{
		UserID:  1,
		Balance: 1000,
	}, {
		UserID:  2,
		Balance: 900,
	}}

	result := DB.Create(&wallets)

	if result.Error != nil {
		panic(result.Error)
	}
}
