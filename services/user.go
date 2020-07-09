package services

import (
	"github.com/johnrazeur/gin-boilerplate/models"
	"github.com/johnrazeur/gin-boilerplate/repositories"
)

// UserService the user service structure
type UserService struct {
}

// Create the user
func (s *UserService) Create(email, username, password string) (*models.User, error) {
	userRepository := new(repositories.UserRepository)
	db := models.DB()
	userRepository.CreateRepository(db)
	var user models.User
	user.Username = username
	user.Email = email
	user.Password = password

	err := userRepository.Create(&user)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

// Login the user
func (s *UserService) Login(email, password string) (*models.User, error) {
	userRepository := new(repositories.UserRepository)
	db := models.DB()
	userRepository.CreateRepository(db)

	user, err := userRepository.Login(email, password)

	if err != nil {
		return user, err
	}

	return user, nil
}
