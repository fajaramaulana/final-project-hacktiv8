package services

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepo
}

func NewPhotoService(photoRepo repositories.PhotoRepo) *PhotoService {
	return &PhotoService{photoRepo: photoRepo}
}

func (s *PhotoService) Create(req *request.CreatePhotoRequest, id int) (map[string]interface{}, error) {
	var photo models.Photo

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl
	photo.UserId = id

	data, err := s.photoRepo.Create(&photo)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":         data.Id,
		"title":      data.Title,
		"caption":    data.Caption,
		"photo_url":  data.PhotoUrl,
		"user_id":    data.UserId,
		"created_at": data.CreatedAt,
	}, nil
}
