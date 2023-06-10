package models

type Programme struct {
	Programme string `gorm:"primaryKey" binding:"required"`
}
