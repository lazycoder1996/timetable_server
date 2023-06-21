package models

import "gorm.io/gorm"

type LoginBody struct {
	gorm.Model
	Reference int `binding:"required" json:"reference"`
	Password  string `binding:"required" json:"password"`
}
