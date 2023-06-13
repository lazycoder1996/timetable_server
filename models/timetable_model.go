package models

type TimetableModel struct {
	CourseName string `json:"course"`
	Course     Course `json:"course_details"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
	Room       Room   `json:"room" gorm:"embedded"`
	Day        string `json:"day"`
}
