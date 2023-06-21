package models

import "github.com/jinzhu/gorm"

type Booking struct {
	gorm.Model
	Reference  int    `json:"reference"`
	Day        string `binding:"required" json:"day"`
	StartTime  int    `binding:"required" json:"start_time"`
	EndTime    int    `binding:"required" json:"end_time"`
	Room       string `binding:"required" json:"room"`
	CourseCode string `binding:"required" json:"course_code"`
	Course     Course `json:"course_details"`
}
