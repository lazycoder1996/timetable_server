package models

type Programme struct {
	Programme string `gorm:"primary_key;not null" json:"programme"`
}
