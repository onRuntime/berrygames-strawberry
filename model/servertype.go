package model

type ServerType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
