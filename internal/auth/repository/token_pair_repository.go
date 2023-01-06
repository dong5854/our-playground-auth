package repository

import "github.com/Team-OurPlayground/our-playground-auth/internal/model"

type TokenPairRepository interface {
	CreateTokenPair(tokePair *model.TokenPair) error            // TODO: 토큰이 생성되면 email, AccessToken, RefreshToken 저장, Refresh 토큰 수명과 일치시키기
	GetTokenPairByEmail(email string) (*model.TokenPair, error) // TODO: RefreshToken 으로 AccessToken 을 새로 발급 받을 때, 두개의 토큰 값의 일치를 찾기 위함
	UpdateTokenPair(email string) error                         // TODO: Access 토큰을 새로 발급 받은 후 업데이트, Refresh 토큰 수명과 일치시키기
}
