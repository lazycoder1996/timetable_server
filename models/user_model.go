package models

import "mime/multipart"

type User struct {
	Reference      int                  `gorm:"PRIMARY_KEY" binding:"required" form:"reference" json:"reference"`
	Password       string               `binding:"required" form:"password"`
	FirstName      string               `binding:"required" form:"firstname"`
	MiddleName     string               `form:"middlename"`
	Surname        string               `binding:"required" form:"surname"`
	Programme      string               `form:"programme_name"`
	Year           int                  `binding:"required" form:"year"`
	PicturePath    multipart.FileHeader `form:"profilepic"`
	ProfilePicture string
	Notification   int `form:"notification"`
	Role           int `gorm:"default:0" form:"role"`
}
