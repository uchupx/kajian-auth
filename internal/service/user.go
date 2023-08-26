package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-auth/internal/dto"
	"github.com/uchupx/kajian-auth/internal/repo"
	"github.com/uchupx/kajian-auth/internal/service/jwt"
)

type UserService struct {
	UserRepo *repo.UserRepo
	JWT      jwt.CryptService
	Redis    *redis.Client
}

func (s *UserService) Login(ctx context.Context, req dto.AuthRequest) (*dto.Response, error) {
	var user dto.User
	model, err := s.UserRepo.FindUserByUsernameEmail(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when find user by username: %w", err)
	}

	user.Model(model)

	token, err := s.JWT.CreateJWTToken(1*time.Hour, user)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when create token: %w", err)
	}

	duration := 1 * time.Hour

	if err := s.Redis.Set(ctx, "redis:auth:token", *token, duration).Err(); err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when set redis: %w", err)
	}

	return &dto.Response{
		Status: 200,
		Data: dto.TokenResponse{
			Token:   *token,
			Expired: int64(duration.Seconds()),
		},
	}, nil
}
