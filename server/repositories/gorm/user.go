package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}
