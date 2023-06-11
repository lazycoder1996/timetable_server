package models

type Room struct {
	RoomName string `gorm:"primaryKey" binding:"required"`
	Size     int    `json:"size"`
}
