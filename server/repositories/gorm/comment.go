package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) Create(comment *models.Comment) (*models.Comment, error) {
	err := r.db.Create(comment).Error
	return comment, err
}

func (r *commentRepo) GetAllComment() ([]models.Comment, error) {
	var comment []models.Comment
	err := r.db.Preload(clause.Associations).Find(&comment).Error
	return comment, err
}

func (r *commentRepo) UpdateCommentById(id int, comment *models.Comment) error {
	return r.db.Where("id = ?", id).Updates(comment).Error
}

func (r *commentRepo) DelteCommentById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.Comment{}).Error
}
