package models

type User struct {
	Reference      int `gorm:"PRIMARY_KEY" binding:"required"`
	Password       string `binding:"required"`
	FirstName      string `binding:"required"`
	MiddleName     string
	Surname        string `binding:"required"`
	Programme      string `json:"programme_name"`
	Year           int    `binding:"required"`
	ProfilePicture string `json:"profilepic"`
	Notification   int    `json:"notification"`
	Role           int    `gorm:"default:0"`
}
