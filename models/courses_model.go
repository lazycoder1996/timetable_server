package models

type Courses struct {
	CourseName string `gorm:"primaryKey" binding:"required"`
}
