package models

import "gorm.io/gorm"

type Schedules struct {
	gorm.Model
	Room               string `gorm:"foreignKey:room_fkey" binding:"required"`
	Programme          string `gorm:"foreignKey:class_fkey" binding:"required"`
	Year               int    `binding:"required"`
	Course             string `gorm:"foreignKey:course_fkey" binding:"required"`
	Day                string `binding:"required"`
	StartTime          int    `binding:"required"`
	EndTime            int    `binding:"required"`
	Recursive          int    `binding:"required"`
	Date               string ``
	NotificationStatus bool   `binding:"required"`
	BookedBy           int    `gorm:"foreignKey:booked_fkey"`
}
