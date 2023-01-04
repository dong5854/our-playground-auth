package entgo

import (
	"context"

	"github.com/Team-OurPlayground/our-playground-auth/ent"
	"github.com/Team-OurPlayground/our-playground-auth/ent/user"
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

func (u *userRepository) FindUserInfoByEmail(email string) (*model.User, error) {
	user, err := u.entClient.User.
		Query().
		Where(user.Email(email)).
		Only(context.TODO())
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrDBInternal, "FindUserInfoByEmail error")
	}

	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsAdmin:   user.IsAdmin,
	}, nil
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
