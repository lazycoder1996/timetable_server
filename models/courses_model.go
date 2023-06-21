package models

type Course struct {
	Code      string `gorm:"primary_key;not null" json:"code"`
	Name      string `gorm:"not null" json:"name"`
}
