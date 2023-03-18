package entity

import "gorm.io/gorm"

type ItemsTrx struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement"`
	ProductId     uint
	Product       Products `gorm:"foreignKey:ProductId"`
	TransactionId uint
	Transaction   Transaction `gorm:"foreignKey:TransactionId"`
}
