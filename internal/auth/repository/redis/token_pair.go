package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository"
	"github.com/Team-OurPlayground/our-playground-auth/internal/model"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

type tokenPairRepository struct {
	redisClient *redis.Client
}

func NewTokenPairRepository(client *redis.Client) repository.TokenPairRepository {
	return &tokenPairRepository{
		redisClient: client,
	}
}

func (t *tokenPairRepository) CreateTokenPair(tokenPair *model.TokenPair) error {
	expiration := tokenPair.ExpiresAt.Sub(time.Now())
	tokenPairByte, err := json.Marshal(tokenPair)
	if err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "json.Marshal error: CreateTokenPair")
	}
	err = t.redisClient.Set(context.TODO(), tokenPair.Email, tokenPairByte, expiration).Err()
	if err != nil {
		return customerror.Wrap(err, customerror.ErrDBInternal, "CreateTokenPair error: CreateTokenPair")
	}
	return nil
}

func (t *tokenPairRepository) GetTokenPairByEmail(email string) (*model.TokenPair, error) {
	resultByte, err := t.redisClient.Get(context.TODO(), email).Bytes()
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrDBInternal, "redis get error: GetTokenPairByEmail")
	}
	result := new(model.TokenPair)
	err = json.Unmarshal(resultByte, result)
	if err != nil {
		return nil, customerror.Wrap(err, customerror.ErrInternalServer, "json.Unmarshal error: GetTokenPairByEmail")
	}
	return result, nil
}

func (t *tokenPairRepository) UpdateTokenPair(tokenPair *model.TokenPair) error {
	expiration := tokenPair.ExpiresAt.Sub(time.Now())
	tokenPairByte, err := json.Marshal(tokenPair)
	if err != nil {
		return customerror.Wrap(err, customerror.ErrInternalServer, "json.Marshal error: UpdateTokenPair")
	}
	err = t.redisClient.Set(context.TODO(), tokenPair.Email, tokenPairByte, expiration).Err()
	if err != nil {
		return customerror.Wrap(err, customerror.ErrDBInternal, "UpdateTokenPair error: UpdateTokenPair")
	}
	return nil
}
