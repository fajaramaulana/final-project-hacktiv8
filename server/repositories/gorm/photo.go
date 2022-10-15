package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repositories.PhotoRepo {
	return &photoRepo{
		db: db,
	}
}

func (r *photoRepo) Create(photo *models.Photo) (*models.Photo, error) {
	err := r.db.Create(photo).Error
	return photo, err
}

func (r *photoRepo) GetAllPhoto() ([]models.Photo, error) {
	var photo []models.Photo
	err := r.db.Preload(clause.Associations).Find(&photo).Error
	return photo, err
}

func (r *photoRepo) UpdatePhotoById(id int, photo *models.Photo) error {
	return r.db.Where("id = ?", id).Updates(photo).Error
}

func (r *photoRepo) DeletePhotoById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.Photo{}).Error
}

func (r *photoRepo) GetPhotoById(id int) (*models.Photo, error) {
	var photo models.Photo
	err := r.db.Where("id = ?", id).First(&photo).Error
	return &photo, err
}
