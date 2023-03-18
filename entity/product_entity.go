package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement"`
	ProductName string
	Price       float32
	Note        string
}
