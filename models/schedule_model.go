package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Room       string `binding:"required" json:"room"`
	Programme  string `binding:"required" json:"programme_name"`
	Year       int    `binding:"required" json:"year"`
	CourseCode string `gorm:"not null" json:"course_code"`
	Course     Course `json:"course_details"`
	Day        string `binding:"required" json:"day"`
	StartTime  int    `binding:"required" json:"start_time"`
	EndTime    int    `binding:"required" json:"end_time"`
	Recursive  int    `binding:"required" json:"recursive"`
	Date       string `json:"date"`
	Status     bool   `binding:"required" json:"status"`
	BookedBy   string `json:"booked_by"`
	BookingID  int    `json:"booking_id"`
}
