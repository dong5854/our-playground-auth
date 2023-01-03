package dto

import "github.com/Team-OurPlayground/our-playground-auth/internal/util/jwt"

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token jwt.Token `json:"token"`
}
