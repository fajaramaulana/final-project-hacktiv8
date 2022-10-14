package services

import (
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(req *request.CreateUserRequest) (map[string]interface{}, error) {
	var user models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Password = string(hash)
	user.Age = req.Age
	user.CreatedAt = time.Now()

	userId, err := s.userRepo.Create(&user)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"age":      user.Age,
		"email":    user.Email,
		"id":       userId,
		"username": user.Username,
	}, nil

}
