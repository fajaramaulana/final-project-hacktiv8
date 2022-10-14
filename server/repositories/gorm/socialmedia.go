package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) repositories.SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (r *socialMediaRepo) Create(socialMedia *models.SocialMedia) error {
	return r.db.Create(socialMedia).Error
}

func (r *socialMediaRepo) GetAllSocialMedia() (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.Find(&socialMedia).Error
	return &socialMedia, err
}

func (r *socialMediaRepo) UpdateSocialMediaById(id int, socialMedia *models.SocialMedia) error {
	return r.db.Where("id = ?", id).Updates(socialMedia).Error
}

func (r *socialMediaRepo) DeleteSocialMediaById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.SocialMedia{}).Error
}

func (r *socialMediaRepo) GetSocialMediaById(id int) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.First(&socialMedia, id).Error
	return &socialMedia, err
}
