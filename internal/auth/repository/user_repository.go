package repository

import "github.com/Team-OurPlayground/our-playground-auth/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserInfoByEmail(email string) (*model.User, error)
	FindUserInfoByID(id int) (*model.User, error)
}
