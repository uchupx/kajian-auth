package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-auth/src/repo"
	"github.com/uchupx/kajian-auth/src/service/jwt"
)

type UserService struct {
	UserRepo         *repo.UserRepo
	ClientRepo       *repo.ClientRepo
	RefreshTokenRepo *repo.RefreshTokenRepo
	JWT              jwt.CryptService
	Redis            *redis.Client
}
