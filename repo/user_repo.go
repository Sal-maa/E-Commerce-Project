package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetIdByName(name string) (entity.User, error)
	Login(name string) (entity.User, error)
	SaveToken(token string) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec("INSERT INTO users(created_at, updated_at, name, email, password, address, phone) VALUES(?,?,?,?,?,?,?)", user.CreatedAt, user.UpdatedAt, user.Name, user.Email, user.Password, user.Address, user.Phone)
	return user, err
}

func (r *userRepository) GetIdByName(name string) (entity.User, error) {
	var user entity.User
	result, err := r.db.Query("SELECT id,name FROM users WHERE name=?", name)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}

	errScan := result.Scan(&user.Id, &user.Name)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}

	if name == user.Name {
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (r *userRepository) Login(name string) (entity.User, error) {
	user := entity.User{}
	result, err := r.db.Query("SELECT name,password FROM users WHERE name=? ", name)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if isExist := result.Next(); !isExist {
		return user, fmt.Errorf("user not exist")
	}

	errScan := result.Scan(&user.Name, &user.Password)

	if errScan != nil {
		fmt.Println(errScan)
		return user, fmt.Errorf("error scanning data")
	}

	if user.Name == name {
		// usernya benar-benar ada
		return user, nil
	}
	// tidak error, tapi usernya tidak ada
	return user, fmt.Errorf("user not found")
}

func (r *userRepository) SaveToken(token string) (string, error) {
	_, err := r.db.Exec("INSERT INTO jwt_token(token) VALUES(?)", token)
	return token, err
}
