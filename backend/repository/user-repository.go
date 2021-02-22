package repository

import (
	"errors"
	"log"
	"shopee-crawler/config"
	"shopee-crawler/model"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var userRepository *userRepo

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// UserRepo interact with db
type UserRepo interface {
	FindByUserAndPassword(username, password string) (*model.User, error)
	FindByID(id int) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// GetUserRepo singleton user repo
func GetUserRepo() UserRepo {
	if userRepository == nil {
		userRepository = &userRepo{
			config.GetDbConnection(),
		}
	}
	return userRepository
}

func (ur *userRepo) FindByUserAndPassword(username, password string) (*model.User, error) {
	var user model.User
	var err error
	if err = userRepository.db.Raw("select * from public.user where username = ? and password = ?", username, password).Take(&user).Error; err != nil {
		err = errors.New("wrong username or password")
	}

	return &user, err
}

// FindByID find user by id
func (ur *userRepo) FindByID(id int) (*model.User, error) {
	var user model.User
	if err := userRepository.db.Where("id = ?", id).Preload("Roles").Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
