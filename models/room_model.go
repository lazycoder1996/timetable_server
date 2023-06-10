package models

type Rooms struct {
	RoomName string `gorm:"primaryKey" binding:"required"`
	Size     int    `json:"size"`
}
