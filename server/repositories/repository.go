package repositories

import "final-project/server/repositories/models"

type UserRepo interface {
	Create(user *models.User) (int, error)
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateById(id int, update *models.User) (*models.User, error)
	Delete(user *models.User) error
}

type PhotoRepo interface {
	Create(photo *models.Photo) error
	GetAllPhoto() (*models.Photo, error)
	UpdatePhotoById(id int, photo *models.Photo) error
	DeletePhotoById(id int) error
}

type CommentRepo interface {
	Create(comment *models.Comment) error
	GetAllComment() (*models.Comment, error)
	UpdateCommentById(id int, comment *models.Comment) error
	DelteCommentById(id int) error
}

type SocialMediaRepo interface {
	Create(socialMedia *models.SocialMedia) error
	GetAllSocialMedia() (*models.SocialMedia, error)
	UpdateSocialMediaById(id int, socialMedia *models.SocialMedia) error
	DeleteSocialMediaById(id int) error
	GetSocialMediaById(id int) (*models.SocialMedia, error)
}
