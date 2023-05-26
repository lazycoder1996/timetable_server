package models

type Users struct {
	UserID int    `gorm:"primaryKey" binding:"required"`
	Class  string `gorm:"foreignKey:class_ref" binding:"required"`
}
