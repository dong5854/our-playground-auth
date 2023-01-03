package service

import "github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"

type AuthService interface {
	SignUp(request *dto.SignUpRequest) error
	SignIn(request *dto.SignInRequest) (bool, error)
	GetToken(email string) (*dto.SignInResponse, error)
}
