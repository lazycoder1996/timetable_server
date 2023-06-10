package models

import "gorm.io/gorm"

type Bookings struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" binding:"required"`
	UserID int    `gorm:"foreignKey:user_fkey" binding:"required"`
	Date   string `binding:"required"`
	Start  int    `binding:"required"`
	End    int    `binding:"required"`
	Room   string `gorm:"foreignKey:room_fkey" binding:"required"`
	Course string `binding:"required"`
}
