package models

type Classes struct {
	ClassName string `gorm:"primaryKey" binding:"required"`
}
