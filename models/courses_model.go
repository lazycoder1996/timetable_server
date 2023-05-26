package models

type Courses struct {
	CourseID string `gorm:"primaryKey" binding:"required"`
}
