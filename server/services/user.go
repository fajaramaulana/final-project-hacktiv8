package services

import (
	"final-project/server/repositories"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

// func (s *UserService) Register(req *request.CreateUserRequest) *UserService {

// }
