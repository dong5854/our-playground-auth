package service

import (
	"encoding/json"
	"time"

	gojwt "github.com/golang-jwt/jwt"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository"
	"github.com/Team-OurPlayground/our-playground-auth/internal/model"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/encrypt"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/jwt"
)

type authServiceImpl struct {
	userRepository      repository.UserRepository
	tokenPairRepository repository.TokenPairRepository
}

func NewAuthService(userRepository repository.UserRepository, tokenPairRepository repository.TokenPairRepository) AuthService {
	return &authServiceImpl{
		userRepository:      userRepository,
		tokenPairRepository: tokenPairRepository,
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
	resp.Token.AccessToken, err = jwt.GenerateAccessToken(jwt.GetPrivateKey(), email)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.GenerateAccessToken error: GetToken")
	}
	resp.Token.RefreshToken, err = jwt.GenerateRefreshToken(jwt.GetPrivateKey(), email)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.GenerateRefreshToken error: GetToken")
	}
	err = a.saveTokenPair(email, resp.Token.AccessToken, resp.Token.RefreshToken)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "saveTokenPair error: GetToken")
	}
	return resp, nil
}

func (a *authServiceImpl) saveTokenPair(email string, accessToken string, refreshToken string) error {
	refreshTokenClaims := new(jwt.CustomClaims)
	refreshTokenByte, err := gojwt.DecodeSegment(refreshToken)
	err = json.Unmarshal(refreshTokenByte, refreshTokenClaims)
	if err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "json.Unmarshal error: saveTokenPair")
	}
	tokenPairModel := &model.TokenPair{
		Email:        email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Unix(refreshTokenClaims.ExpiresAt, 0),
	}
	err = a.tokenPairRepository.CreateTokenPair(tokenPairModel)
	if err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "tokenPairRepository.CreateTokenPair: saveTokenPair")
	}
	return err
}

func (a *authServiceImpl) Refresh(request *dto.RefreshRequest) (*dto.RefreshResponse, error) {
	token, err := a.verifyRefresh(request.AccessToken, request.RefreshToken)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.VerifyToken error: Refresh")
	}
	resp := new(dto.RefreshResponse)
	resp.Token.AccessToken, err = jwt.GenerateAccessToken(jwt.GetPrivateKey(), token.Email)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.GenerateAccessToken error: Refresh")
	}
	resp.Token.RefreshToken, err = jwt.GenerateRefreshToken(jwt.GetPrivateKey(), token.Email)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.GenerateRefreshToken error: Refresh")
	}
	err = a.saveTokenPair(token.Email, resp.Token.AccessToken, resp.Token.RefreshToken)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "authServiceImpl.saveTokenPair error: Refresh")
	}
	return resp, nil
}

func (a *authServiceImpl) verifyRefresh(accessToken string, refreshToken string) (*jwt.CustomClaims, error) {
	token, err := jwt.VerifyToken(refreshToken)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "jwt.VerifyToken error: verifyRefreshToken")
	}
	tokenPairModel, err := a.tokenPairRepository.GetTokenPairByEmail(token.Email)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrDBInternal, "tokenPairRepository.GetTokenPairByEmail error: verifyRefreshToken")
	}
	if tokenPairModel.AccessToken != accessToken {
		return nil, customerror.New(customerror.ErrInternalServer, "accessToken doesn't match")
	}
	if tokenPairModel.RefreshToken != refreshToken {
		return nil, customerror.New(customerror.ErrInternalServer, "refreshToken doesn't match")
	}
	return token, err
}
