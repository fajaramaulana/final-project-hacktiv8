package services

import (
	"final-project/server/controllers/view"
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
	"time"
)

type CommentService struct {
	commentRepo repositories.CommentRepo
}

func NewCommentService(commentRepo repositories.CommentRepo) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) Create(idUser int, req *request.CreateCommentRequest) (view.ResponseCreateComment, error) {
	var comment models.Comment

	comment.Message = req.Message
	comment.PhotoId = req.PhotoId
	comment.UserId = idUser
	comment.CreatedAt = time.Now()
	data, err := s.commentRepo.Create(&comment)

	if err != nil {
		return view.ResponseCreateComment{}, err
	}

	return view.ResponseCreateComment{
		Id:        data.Id,
		Message:   data.Message,
		PhotoId:   data.PhotoId,
		UserId:    data.UserId,
		CreatedAt: data.CreatedAt,
	}, nil
}

func (s *CommentService) GetAll() ([]view.ResponseGetAllComment, error) {
	data, err := s.commentRepo.GetAllComment()

	if err != nil {
		return []view.ResponseGetAllComment{}, err
	}

	var response []view.ResponseGetAllComment

	for _, v := range data {
		response = append(response, view.ResponseGetAllComment{
			Id:        v.Id,
			Message:   v.Message,
			PhotoId:   v.PhotoId,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: view.ResponseWithUserIdComment{
				Id:       v.User.Id,
				Username: v.User.Username,
				Email:    v.User.Email,
			},
			Photo: view.ResponseWithPhotoIdComment{
				Id:       v.Photo.Id,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserId:   v.Photo.UserId,
			},
		})
	}

	return response, nil
}
