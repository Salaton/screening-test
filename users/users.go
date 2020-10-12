package users

// This package will handle our create user & login functionalities..

import (
	"log"

	model "github.com/Salaton/screening-test/graph/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserMethodsInterface to hold methods that interact with the users
type UserMethodsInterface interface {
	GetUserID(username string) (int, error)
	CreateUser(model.CreatedUser)
}

// PostgresClient for db references
type DatabaseClient struct {
	db *gorm.DB
}

func (ps *DatabaseClient) CreateUser(user model.CreatedUser) {
	if err := ps.db.Create(&model.User{
		Username: user.Username,
		Password: HashPassword(user.Password),
	}).Error; err != nil {
		log.Printf("Something went wrong %v", err)
	}
}

// HashPassword to hash the password since we dont want to store
// a password string in the database
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckPasswordHash that compares the password and generated hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserID -> Check if a user exists in a DB using the username
func (ps *DatabaseClient) GetUserID(username string) (int, error) {
	var user model.User
	row := ps.db.Table("users").Where("username = ?", username).Select("id").Row()
	var id int
	row.Scan(&user.ID)

	return id, nil
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
