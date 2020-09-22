package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	PlayerName      string `gorm:"unique"`
	Nickname        string

	Ranks           []Rank `gorm:"many2many:player_ranks;"`
	PermissionLevel int

	Coins           int
	Credits         int

	Experience      int

	Locale          string
}
