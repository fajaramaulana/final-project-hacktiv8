package gorm

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"

	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepo) GetAllComment() (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Find(&comment).Error
	return &comment, err
}

func (r *commentRepo) UpdateCommentById(id int, comment *models.Comment) error {
	return r.db.Where("id = ?", id).Updates(comment).Error
}

func (r *commentRepo) DelteCommentById(id int) error {
	return r.db.Where("id = ?", id).Delete(&models.Comment{}).Error
}
