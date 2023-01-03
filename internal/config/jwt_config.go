package config

import (
	"crypto/rsa"
	"sync"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/jwt"
)

var (
	JWTOnce    sync.Once
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func InitJWTKeys() {
	JWTOnce.Do(func() {
		var err error
		privateKey, publicKey, err = jwt.InitRSAKey()
		if err != nil {
			panic(customerror.Wrap(err, customerror.ErrInternalServer, "config InitJWTKeys error"))
		}
	})
}

func GetPrivateKey() *rsa.PrivateKey {
	InitJWTKeys()
	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	InitJWTKeys()
	return publicKey
}
