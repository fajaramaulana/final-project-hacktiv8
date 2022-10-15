package services

import (
	"errors"
	"final-project/server/controllers/view"
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
	"fmt"
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

func (s *CommentService) Update(idUser int, idComment int, req *request.UpdateCommentRequest) (view.ResponseUpdateComment, error) {
	var comment models.Comment

	comment.Message = req.Message

	checkIfExist, err := s.commentRepo.CheckCommentByIdAndUserId(idComment, idUser)

	if !checkIfExist {
		return view.ResponseUpdateComment{}, errors.New("Unauthorized")
	}

	data, err := s.commentRepo.UpdateCommentById(idComment, &comment)

	fmt.Printf("%# v", data)

	if err != nil {
		return view.ResponseUpdateComment{}, err
	}

	return view.ResponseUpdateComment{
		Id:        data.Id,
		Title:     data.Photo.Title,
		Caption:   data.Photo.Caption,
		PhotoUrl:  data.Photo.PhotoUrl,
		UserId:    data.UserId,
		UpdatedAt: data.UpdatedAt,
	}, nil
}
