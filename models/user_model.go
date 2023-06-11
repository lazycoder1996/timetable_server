package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Reference      string `gorm:"primaryKey" binding:"required"`
	Password       string `binding:"required"`
	FirstName      string `binding:"required"`
	MiddleName     string ``
	Surname        string `binding:"required"`
	Programme      string `gorm:"foreignKey:programme_ref" binding:"required"`
	Year           int    `binding:"required"`
	ProfilePicture string `json:"profilepic"`
	Notification   int `json:"notification"`
	Role           int    `gorm:"default:0"`
}
