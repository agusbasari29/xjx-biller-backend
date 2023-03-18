package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserId      uint
	User        Users `gorm:"foreignKey:UserId"`
	TotalAmount float32
	Note        string
}
