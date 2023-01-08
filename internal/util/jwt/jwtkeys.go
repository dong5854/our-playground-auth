package jwt

import (
	"crypto/rsa"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang-jwt/jwt"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

var (
	once          sync.Once
	privateKey    *rsa.PrivateKey
	rawPrivateKey []byte
	publicKey     *rsa.PublicKey
	rawPublicKey  []byte
)

func InitJWTKeys() {
	once.Do(func() {
		currentDIR, err := os.Getwd()
		if err != nil {
			log.Fatal(customerror.Wrap(err, customerror.ErrInternalServer, "os.Getwd() error : InitJWTKeys"))
		}

		var (
			privateKeyPath = filepath.Join(currentDIR, "private.key")
			publicKeyPath  = filepath.Join(currentDIR, "public.pem")
		)

		rawPrivateKey, err = os.ReadFile(privateKeyPath)
		if err != nil {
			log.Fatal(customerror.Wrap(err, customerror.ErrInternalServer, "os.ReadFile(privateKeyPath) error : InitJWTKeys"))
		}

		rawPublicKey, err = os.ReadFile(publicKeyPath)
		if err != nil {
			log.Fatal(customerror.Wrap(err, customerror.ErrInternalServer, "os.ReadFile(publicKeyPath) error : InitJWTKeys"))
		}

		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(rawPrivateKey)
		if err != nil {
			log.Fatal(customerror.Wrap(err, customerror.ErrInternalServer, "jwt.ParseRSAPrivateKeyFromPEM(rawPrivateKey) error : InitJWTKeys"))
		}

		publicKey, err = jwt.ParseRSAPublicKeyFromPEM(rawPublicKey)
		if err != nil {
			log.Fatal(customerror.Wrap(err, customerror.ErrInternalServer, "jwt.ParseRSAPublicKeyFromPEM(rawPublicKey) error : InitJWTKeys"))
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

func GetRawPublicKey() []byte {
	InitJWTKeys()
	return rawPublicKey
}
