package model

import "gorm.io/gorm"

type Server struct {
	gorm.Model

	ServerTypeID int
	ServerType   ServerType `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Players      []Player   `gorm:"many2many:server_players;"`
}
