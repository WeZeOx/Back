package models

import (
	"time"
)

type User struct {
	ID             string `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	Username       string `json:"username" gorm:"unique"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
	Email          string `json:"email" gorm:"unique"`
}
