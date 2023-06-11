package models

type Course struct {
	CourseName string `gorm:"primaryKey" binding:"required"`
}
