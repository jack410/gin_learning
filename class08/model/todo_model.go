package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
}
