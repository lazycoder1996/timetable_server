package models

import "gorm.io/gorm"

type LoginBody struct {
	gorm.Model
	Reference string `binding:"required" json:"reference"`
	Password  string `binding:"required" json:"password"`
}
