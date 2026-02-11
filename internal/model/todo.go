package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"title"`
	Completed bool   `gorm:"completed"`
	UserID    uint   `gorm:"column:user_id"`
}
