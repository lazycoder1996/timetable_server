package models

type Schedule struct {
	RoomName   string `json:"room" gorm:"primary_key;not null"`
	Room       Room   `json:"room_name"`
	Programme  string `binding:"required" json:"programme"`
	Year       int    `binding:"required" json:"year"`
	CourseCode string `gorm:"not null" json:"course_code"`
	Course     Course `json:"course_details"`
	Day        string `binding:"required" json:"day" gorm:"primary_key"`
	StartTime  int    `binding:"required" json:"start_time" gorm:"primary_key"`
	EndTime    int    `binding:"required" json:"end_time" gorm:"primary_key"`
	Recursive  bool   `json:"recursive"`
	Date       string `json:"date"`
	Status     bool   `binding:"required" json:"status" gorm:"default:true"`
	BookedBy   int    `json:"booked_by"`
	BookingID  int    `json:"booking_id"`
}
