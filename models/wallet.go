package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	PlayerID uint
	Player   Player
	Balance  float64
}
