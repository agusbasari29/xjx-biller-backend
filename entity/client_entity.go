package entity

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Client   string
	User     string
	Password string
	Status   string
}
