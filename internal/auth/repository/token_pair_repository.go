package repository

import "github.com/Team-OurPlayground/our-playground-auth/internal/model"

type TokenPairRepository interface {
	CreateTokenPair(tokePair *model.TokenPair) error
	GetTokenPairByEmail(email string) (*model.TokenPair, error)
	UpdateTokenPair(tokenPair *model.TokenPair) error
}
