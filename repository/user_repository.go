package repository

import (
	"github.com/agusbasari29/xjx-biller-backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.Users) (entity.Users, error)
	GetByUsername(username string) interface{}
	GetUser(user entity.Users) (entity.Users, error)
	UserIsExist(username string) bool
	UpdateUser(user entity.Users) (entity.Users, error)
	UpdateUserPassword(user entity.Users) error
	DeleteUser(user entity.Users) error
}

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

func (r *userRepository) GetByUsername(username string) interface{} {
	user := entity.Users{}
	res := r.db.Raw("SELECT * FROM users WHERE username=@Username", map[string]interface{}{"Username": username}).Take(&user)
	if res.Error != nil {
		return user
	}
	return nil
}

func (r *userRepository) GetUser(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("SELECT * FROM users WHERE id=@ID", user).Take(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) UserIsExist(username string) bool {
	var user entity.Users
	res := r.db.Raw("SELECT * FROM users WHERE username=@Username", map[string]interface{}{"Username": username}).Take(&user)
	return res.Error == nil
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

func (r *userRepository) DeleteUser(user entity.Users) error {
	err := r.db.Raw("DELETE FROM users WHERE id=@ID", user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
