package jwt

import (
	"crypto/rsa"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

const (
	refreshAudience = "refresh"
	accessAudience  = "access"
)

func InitRSAKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	currentDIR, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	var (
		privateKey     *rsa.PrivateKey
		publicKey      *rsa.PublicKey
		privateKeyPath = filepath.Join(currentDIR, "private.key")
		publicKeyPath  = filepath.Join(currentDIR, "public.pem")
	)

	rawPrivateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, nil, err
	}

	rawPublicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(rawPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, publicKey, nil
}

func GenerateAccessToken(privateKey *rsa.PrivateKey, email string) (string, error) {
	accessTokenClaim := &CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // expires within 15 minutes.
			Audience:  accessAudience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaim)

	jwtToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", customerror.Wrap(err, customerror.ErrInternalServer, "generateAccessToken error")
	}

	return jwtToken, nil
}

func GenerateRefreshToken(privateKey *rsa.PrivateKey, email string) (string, error) {
	refreshTokenClaim := &CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // expires within a week.
			Audience:  refreshAudience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaim)

	jwtToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", customerror.Wrap(err, customerror.ErrInternalServer, "generateRefreshToken error")
	}

	return jwtToken, nil
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, new(CustomClaims), func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, customerror.New(customerror.ErrInternalServer, "Unexpected signing method: %v", token.Header["alg"])
		}
		return GetPublicKey(), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*CustomClaims), nil
}
