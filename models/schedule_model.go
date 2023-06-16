package models

type Schedule struct {
	RoomName   string `json:"room_name" gorm:"primary_key;not null"`
	Room       Room   `json:"room_details"`
	Programme  string `binding:"required" json:"programme_name"`
	Year       int    `binding:"required" json:"year"`
	CourseCode string `gorm:"not null" json:"course_code"`
	Course     Course `json:"course_details"`
	Day        string `binding:"required" json:"day" gorm:"primary_key"`
	StartTime  int    `binding:"required" json:"start_time" gorm:"primary_key"`
	EndTime    int    `binding:"required" json:"end_time" gorm:"primary_key"`
	Recursive  int    `binding:"required" json:"recursive"`
	Date       string `json:"date"`
	Status     bool   `binding:"required" json:"status"`
	BookedBy   string `json:"booked_by"`
	BookingID  int    `json:"booking_id"`
}
