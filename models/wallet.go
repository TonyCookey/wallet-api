package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserID  uint
	User    User
	Balance float64
}
