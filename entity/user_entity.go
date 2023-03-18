package entity

import "gorm.io/gorm"

type UserRole string

const (
	Admin UserRole = "admin"
	User  UserRole = "user"
)

type Users struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Password string
	Role     UserRole
}
