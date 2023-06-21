package models

type Room struct {
	Name     string `gorm:"primary_key;not null" json:"room_name"`
	Location string `gorm:"default:'Engineering'" json:"location"`
	Type     string `json:"type" gorm:"default:'class'"`
	Size     int    `json:"size" gorm:"default:0"`
}

type RoomStatusResponse struct {
	Room      Room   `json:"room"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
	Course    Course `json:"course"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	Status    bool   `json:"status"`
}

type RoomAvailableTimes struct {
	Day       string `json:"day"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}
