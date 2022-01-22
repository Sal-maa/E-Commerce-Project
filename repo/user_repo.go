package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type UserRepository interface {
	CheckUser(userChecked entity.CreateUserRequest) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	GetIdByName(name string) (entity.User, error)
	Login(name string) (entity.User, error)
	SaveToken(token string) (string, error)
	GetUser(idParam int) (entity.User, error)
	DeleteUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CheckUser(userChecked entity.CreateUserRequest) (entity.User, error) {
	user := entity.User{}
	result, err := r.db.Query("SELECT username FROM users WHERE username=?", userChecked.Username)
	if err != nil {
		return user, err
	}
	defer result.Close()
	if isExist := result.Next(); isExist {
		return user, fmt.Errorf("user already exist")
	}

	if user.Username != userChecked.Username {
		// usernya belum ada
		return user, nil
	}

	return user, nil
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec(`INSERT INTO 
								users(created_at, updated_at, username, email, password, address, phone) 
						VALUES(?,?,?,?,?,?,?)`, user.CreatedAt, user.UpdatedAt, user.Username, user.Email, user.Password, user.Address, user.Phone)
	return user, err
}

func (r *userRepository) GetIdByName(name string) (entity.User, error) {
	var user entity.User
	result, err := r.db.Query("SELECT id, username FROM users WHERE username=?", name)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}

	errScan := result.Scan(&user.Id, &user.Username)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}

	if name == user.Username {
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (r *userRepository) Login(name string) (entity.User, error) {
	user := entity.User{}
	result, err := r.db.Query("SELECT username,password FROM users WHERE username=? ", name)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if isExist := result.Next(); !isExist {
		return user, fmt.Errorf("user not exist")
	}

	errScan := result.Scan(&user.Username, &user.Password)

	if errScan != nil {
		fmt.Println(errScan)
		return user, fmt.Errorf("error scanning data")
	}

	if user.Username == name {
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

func (r *userRepository) GetUser(idParam int) (entity.User, error) {
	var user entity.User
	result, err := r.db.Query("SELECT id, username, email, password, address, phone FROM users WHERE id=?", idParam)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}

	errScan := result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address, &user.Phone)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}

	if idParam == user.Id {
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

// func (r *userRepository) DeleteUser(user entity.User) (entity.User, error) {
// 	_, err := r.db.Exec("UPDATE users SET deleted_at=? WHERE id=?", user.DeletedAt, user.Id)
// 	return user, err
// }

func (r *userRepository) DeleteUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec("DELETE FROM products WHERE id =", user.Id)
	return user, err
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec(`UPDATE users 
						SET updated_at = ?, username = ?, email = ?, password = ?, address = ?, phone = ? WHERE id = ?`, user.UpdatedAt, user.Username, user.Email, user.Password, user.Address, user.Phone, user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
