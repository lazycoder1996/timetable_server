package models

type Room struct {
	RoomName string `gorm:"primary_key;not null" binding:"required" json:"room_name"`
	Size     int    `json:"size"`
}
