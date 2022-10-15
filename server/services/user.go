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

func (s *UserService) Login(req *request.UserLoginRequest) (string, error) {
	data, err := s.userRepo.FindByEmail(req.Email)

	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))

	if err != nil {
		return "", err
	}

	return req.Email, nil
}

func (s *UserService) Update(id int, req *request.UpdateUserRequest) (map[string]interface{}, error) {
	var user models.User
	user.Username = req.Username
	user.Email = req.Email

	data, err := s.userRepo.UpdateById(id, &user)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":         data.Id,
		"email":      data.Email,
		"username":   data.Username,
		"age":        data.Age,
		"updated_at": data.UpdatedAt,
	}, nil
}

func (s *UserService) Delete(email string) (map[string]interface{}, error) {
	err := s.userRepo.DeleteByEmail(email)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Your account has been successfully deleted",
	}, nil
}

func (s *UserService) GetUserIdByEmail(email string) (int, error) {
	data, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return 0, err
	}

	return data.Id, nil
}
