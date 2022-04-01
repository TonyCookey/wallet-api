package services

import (
	"wallet-api/database"
	"wallet-api/models"
	"wallet-api/utils"
)

func ValidateUser(email string, password string) bool {
	player := models.Player{}
	err := database.DB.Where("email_address = ?", email).Find(&player).Error
	if err != nil {
		return false
	}
	// compare password with the hashed password
	isValidPassword := utils.CheckPasswordHash(password, player.Password)

	if !isValidPassword {
		return false
	}
	return true

}
