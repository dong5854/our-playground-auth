package service

import (
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository"
	"github.com/Team-OurPlayground/our-playground-auth/internal/model"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

type authServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authServiceImpl{
		userRepository: userRepository,
	}
}

func (a *authServiceImpl) SignUp(request *dto.SignUpRequest) error {
	user := &model.User{
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		IsAdmin:   false,
	}
	if err := a.userRepository.CreateUser(user); err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "userRepository.CreateUser error at authService SignUp")
	}
	return nil
}

func (a *authServiceImpl) SignIn(request *dto.SignInRequest) error {
	return nil
}
