package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-api/pkg/errors"
	"github.com/uchupx/kajian-auth/internal/dto"
	"github.com/uchupx/kajian-auth/internal/repo"
	"github.com/uchupx/kajian-auth/internal/service/jwt"
	"github.com/uchupx/kajian-auth/pkg"
	"github.com/uchupx/kajian-auth/pkg/enums"
)

type UserService struct {
	UserRepo   *repo.UserRepo
	ClientRepo *repo.ClientRepo
	JWT        jwt.CryptService
	Redis      *redis.Client
}

func (s *UserService) Login(ctx context.Context, req dto.AuthRequest) (*dto.Response, error) {
	var user dto.User
	if req.Username == nil || req.Password == nil {
		return nil, fmt.Errorf("[UserService - Login] username or password is required")
	}

	client, err := s.ClientRepo.FindAppsByKey(ctx, req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when find client app: %w", err)
	}

	isClientValid, err := s.JWT.Verify(req.ClientSecret, client.Secret.String)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when verify client secret: %w", err)
	}

	if !isClientValid {
		return nil, errors.ErrUnauthorized
	}

	model, err := s.UserRepo.FindUserByUsernameEmail(ctx, *req.Username)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when find user by username: %w", err)
	}

	// sign, err := s.JWT.CreateSignPSS(req.Password)
	// if err != nil {
	// 	return nil, fmt.Errorf("[UserService - Login] error when create signature password: %w", err)
	// }

	isValid, err := s.JWT.Verify(*req.Password, model.Password.String)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when verify value: %w", err)
	}

	if !isValid {
		return nil, errors.ErrUnauthorized
	}

	user.Model(model)

	token, err := s.JWT.CreateAccessToken(1*time.Hour, user)
	if err != nil {
		return nil, fmt.Errorf("[UserService - Login] error when create access token: %w", err)
	}

	duration := 1 * time.Hour

	if err := s.Redis.Set(ctx, fmt.Sprintf(enums.RedisKeyAuthorizationToken, strings.Split(*token, ".")[2]), pkg.JsonStringify(user), duration).Err(); err != nil {
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

func (s *UserService) SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.Response, error) {
	client, err := s.ClientRepo.FindAppsByKey(ctx, req.ClientKey)
	if err != nil {
		return nil, fmt.Errorf("[UserService - SignUp] error when find client app: %w", err)
	} else if client == nil {
		return nil, errors.ErrNotFound
	}

	signPassword, err := s.JWT.CreateSignPSS(req.Password)
	if err != nil {
		return nil, fmt.Errorf("[UserService - SignUp] error when create signature password: %w", err)
	}

	now := time.Now()

	newUser := dto.User{
		Username:    req.Username,
		Password:    signPassword,
		Email:       req.Email,
		ClientAppId: client.ID.String,
		Created:     now,
	}

	id, err := s.UserRepo.Insert(ctx, newUser.ToModel())
	if err != nil {
		return nil, fmt.Errorf("[UserService - SignUp] error when creating user: %w", err)
	}

	newUser.ID = *id

	return &dto.Response{
		Status: 201,
		Data: dto.EntityResponse{
			Id:     *id,
			Entity: "users",
		},
	}, nil
}

func (s *UserService) RetrieveUser(ctx context.Context, token string) (*dto.User, error) {
	resToken, err := s.JWT.VerifyJWTToken(token)
	if err != nil {
		return nil, fmt.Errorf("[UserService - RetrieveUser] error when verify token: %w", err)
	}

	bytes, err := json.Marshal(resToken)
	if err != nil {
		return nil, fmt.Errorf("[UserService - RetrieveUser] error when marshal token: %w", err)
	}

	var user dto.User

	if err := json.Unmarshal(bytes, &user); err != nil {
		return nil, fmt.Errorf("[UserService - RetrieveUser] error when unmarshal token: %w", err)
	}

	return &user, nil
}

func (s *UserService) AddClient(ctx context.Context, req dto.ClientPost) (*dto.Response, error) {

	secret := RandomString(20)
	clientSecret, err := s.JWT.CreateSignPSS(secret)
	if err != nil {
		return nil, fmt.Errorf("[UserService - AddClient] error when create signature password: %w", err)
	}

	data := dto.Client{
		Name:   req.Name,
		Key:    RandomString(20),
		Secret: clientSecret,
	}

	id, err := s.ClientRepo.Insert(ctx, data.ToModel())
	if err != nil {
		return nil, fmt.Errorf("[UserService - AddClient] error when insert client: %w", err)
	}

	return &dto.Response{
		Status: 201,
		Data: dto.EntityResponse{
			Id:     id,
			Entity: "client_apps",
			Meta: map[string]interface{}{
				"secret": secret,
			},
		},
	}, nil
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
