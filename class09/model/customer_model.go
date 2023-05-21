package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserID      int    `json:"user_id"`
}
