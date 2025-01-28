package repositories

import (
	"errors"
	"gorm.io/gorm"
	"instashop/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (a *UserRepository) Create(input *models.User) error {
	return a.db.Create(input).Error
}

func (a *UserRepository) FetchOne(filter models.User) (*models.User, bool, error) {
	var user models.User
	if err := a.db.Where(filter).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &user, true, nil
}
