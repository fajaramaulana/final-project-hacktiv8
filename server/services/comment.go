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
