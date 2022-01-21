package service

import (
	"fmt"
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
	GetUserByIdService(id int) (entity.User, error)
	DeleteUserService(id int) (entity.User, error)
	UpdateUserService(id int, user entity.EditUserRequest) (entity.User, error)
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

func (s *userService) GetUserByIdService(id int) (entity.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (s *userService) DeleteUserService(id int) (entity.User, error) {
	userID, err := s.GetUserByIdService(id)
	if err != nil {
		return userID, err
	}

	userID.DeletedAt = time.Now()
	deleteUser, err := s.repository.DeleteUser(userID)
	return deleteUser, err
}

func (s *userService) UpdateUserService(id int, userUpdate entity.EditUserRequest) (entity.User, error){
	user, err := s.GetUserByIdService(id)
	if err != nil {
		return user, err
	}
	
	//user := entity.User{}
	user.UpdatedAt = time.Now()
	user.Address = userUpdate.Address
	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	user.Phone = userUpdate.Phone
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.MinCost)
	
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	updateUser, err := s.repository.UpdateUser(user)
	
	fmt.Println(user)
	
	return updateUser, err 
}