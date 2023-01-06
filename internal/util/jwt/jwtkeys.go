package jwt

import (
	"crypto/rsa"
	"sync"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

var (
	once       sync.Once
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func InitJWTKeys() {
	once.Do(func() {
		var err error
		privateKey, publicKey, err = InitRSAKey()
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
