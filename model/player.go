package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name string `gorm:"unique"`
}
