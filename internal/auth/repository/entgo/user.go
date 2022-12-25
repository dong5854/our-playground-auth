package entgo

import (
	"context"

	"github.com/Team-OurPlayground/our-playground-auth/ent"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository"
	"github.com/Team-OurPlayground/our-playground-auth/internal/model"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

type userRepository struct {
	entClient *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{
		entClient: client,
	}
}

func (u *userRepository) CreateUser(user *model.User) error {
	_, err := u.entClient.User.
		Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetUserName(user.UserName).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		Save(context.TODO())
	if err != nil {
		return customerror.Wrap(err, customerror.ErrDBInternal, "CreateUser error")
	}
	return nil
}

func (u *userRepository) FindUserInfoByID(id int) (*model.User, error) {
	result, err := u.entClient.User.Get(context.TODO(), id)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "FindUserInfoByID error")
	}

	user := &model.User{
		ID:        result.ID,
		Email:     result.Email,
		Password:  result.Password,
		UserName:  result.UserName,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		IsAdmin:   result.IsAdmin,
	}
	return user, nil
}
