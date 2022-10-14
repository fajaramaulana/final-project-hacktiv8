package services

import (
	"final-project/server/repositories"
	"final-project/server/request"
	"fmt"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(req *request.CreateUserRequest) (map[string]interface{}, error) {
	fmt.Println(req)
	return nil
}
