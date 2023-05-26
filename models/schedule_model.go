package models

type Schedules struct {
	ScheduleID         int    `gorm:"primaryKey"`
	Room               string `gorm:"foreignKey:room_fkey"`
	Class              string `gorm:"foreignKey:class_fkey"`
	Course             string `gorm:"foreignKey:course_fkey"`
	Day                string `binding:"required"`
	StartTime          string `binding:"required"`
	EndTime            string `binding:"required"`
	Recursive          bool   `binding:"required"`
	Date               string `binding:"required"`
	NotificationStatus bool   `binding:"required"`
	BookedBy           int    `gorm:"foreignKey:booked_fkey"`
}
