package models

import "gorm.io/gorm"

type LoginBody struct {
	gorm.Model
	Reference string `binding:"required"`
	Password  string `binding:"required"`
}
