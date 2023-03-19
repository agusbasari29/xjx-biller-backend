package repository

import (
	"github.com/agusbasari29/xjx-biller-backend/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.Users) (entity.Users, error) {
	if user.Role == "" {
		user.Role = entity.User
	}
	err := r.db.Raw("INSERT INTO users (username, password, role) VALUES (@Username, @Password, @Role)", user).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("UPDATE users SET username=@Username, role=@Role WHERE id=@ID", user).Save(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserPassword(user entity.Users) error {
	err := r.db.Raw("UPDATE users SET password=@Password", user).Save(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *userRepository) DeleteUser()
