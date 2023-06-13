package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	Day       string `binding:"required"`
	Date      string `binding:"required"`
	Start     int    `binding:"required"`
	End       int    `binding:"required"`
	Room      string `binding:"required"`
	Course    string `binding:"required"`
	Programme string `binding:"required"`
	Year      int    `binding:"required"`
}
