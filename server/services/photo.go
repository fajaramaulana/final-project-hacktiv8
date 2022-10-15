package services

import (
	"final-project/server/controllers/view"
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

func (s *PhotoService) Create(req *request.CreatePhotoRequest, id int) (view.ResponseCreatePhoto, error) {
	var photo models.Photo

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl
	photo.UserId = id

	data, err := s.photoRepo.Create(&photo)

	if err != nil {
		return view.ResponseCreatePhoto{}, err
	}

	return view.ResponseCreatePhoto{
		Id:        data.Id,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserId:    data.UserId,
		CreatedAt: data.CreatedAt,
	}, nil
}

func (s *PhotoService) GetAll() ([]view.ResponseGetAllPhoto, error) {
	data, err := s.photoRepo.GetAllPhoto()

	if err != nil {
		return []view.ResponseGetAllPhoto{}, err
	}

	var response []view.ResponseGetAllPhoto

	for _, v := range data {
		response = append(response, view.ResponseGetAllPhoto{
			Id:        v.Id,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: view.ResponseWithUserId{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return response, nil

}

func (s *PhotoService) Update(req *request.UpdatePhotoRequest, id int, idUser int) (view.ResponseUpdatePhoto, error) {
	var photo models.Photo

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl
	photo.UserId = idUser

	data, err := s.photoRepo.UpdatePhotoByIdAndUserId(id, idUser, &photo)

	if err != nil {
		return view.ResponseUpdatePhoto{}, err
	}

	return view.ResponseUpdatePhoto{
		Id:        data.Id,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserId:    data.UserId,
		UpdatedAt: data.UpdatedAt,
	}, nil

}
