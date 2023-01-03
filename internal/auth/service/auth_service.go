package service

import (
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository"
	"github.com/Team-OurPlayground/our-playground-auth/internal/config"
	"github.com/Team-OurPlayground/our-playground-auth/internal/model"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/encrypt"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/jwt"
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
		Password:  encrypt.Sha256(request.Password),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		IsAdmin:   false,
	}
	if err := a.userRepository.CreateUser(user); err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "userRepository.CreateUser error at authService SignUp")
	}
	return nil
}

func (a *authServiceImpl) SignIn(request *dto.SignInRequest) (bool, error) {
	user, err := a.userRepository.FindUserInfoByEmail(request.Email)
	if err != nil {
		return false, err
	}
	return validateUser(user, request.Password), nil
}

func validateUser(user *model.User, password string) bool {
	if user.Password != encrypt.Sha256(password) {
		return false
	}
	return true
}

func (a *authServiceImpl) GetToken(email string) (*dto.SignInResponse, error) {
	resp := new(dto.SignInResponse)
	var err error
	resp.Token.AccessToken, err = jwt.GenerateAccessToken(config.GetPrivateKey(), email)
	resp.Token.RefreshToken, err = jwt.GenerateRefreshToken(config.GetPrivateKey(), email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
