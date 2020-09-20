package model

import "gorm.io/gorm"

type Server struct {
	gorm.Model

	ServerTypeID int
	ServerType   ServerType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Players      []Player   `gorm:"foreignkey:Name"`
}
