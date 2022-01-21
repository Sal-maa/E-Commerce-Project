package service

import (
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUserService(userCreate entity.CreateUserRequest) (entity.User, error)
	GetUserByNameService(name string) (entity.User, error)
	LoginUserService(login entity.LoginUserRequest) (entity.User, error)
	SaveTokenService(token string) (string, error)
}

type userService struct {
	repository repo.UserRepository
}

func NewUserService(repository repo.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUserService(userCreate entity.CreateUserRequest) (entity.User, error) {
	user := entity.User{}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Name = userCreate.Name
	user.Email = userCreate.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Address = userCreate.Address
	user.Phone = userCreate.Phone

	createUser, err := s.repository.CreateUser(user)
	return createUser, err
}

func (s *userService) GetUserByNameService(name string) (entity.User, error) {
	user, err := s.repository.GetIdByName(name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) LoginUserService(login entity.LoginUserRequest) (entity.User, error) {
	name := login.Name
	password := login.Password

	user, err := s.repository.Login(name)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) SaveTokenService(token string) (string, error) {
	sToken, err := s.repository.SaveToken(token)
	return sToken, err
}
