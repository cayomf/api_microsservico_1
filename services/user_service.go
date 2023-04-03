package services

import (
	"errors"

	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	Login(email, password string) (string, error)
}

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) bool
}

type userService struct {
	userRepository repositories.UserRepository
	jwtService     JWTService
}

func NewUserService(userRepository repositories.UserRepository, jwtService JWTService) UserService {
	return &userService{userRepository: userRepository, jwtService: jwtService}
}

func (us *userService) CreateUser(user *models.User) error {
	user.Password = SHA256Encoder(user.Password)
	return us.userRepository.CreateUser(user)
}

func (us *userService) Login(email, password string) (string, error) {
	user, err := us.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("cannot find user")
	}

	if user.Password != SHA256Encoder(password) {
		return "", errors.New("invalid credentials")
	}

	token, err := us.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
