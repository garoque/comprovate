package user

import (
	"github.com/garoque/comprovate/model"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) Database {
	return &user{db}
}

func (u *user) FindByEmail(email string) (*model.User, error) {
	var user model.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
