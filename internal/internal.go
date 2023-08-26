package internal

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-api/pkg/mysql"
	kajianRedis "github.com/uchupx/kajian-api/pkg/redis"
	"github.com/uchupx/kajian-auth/config"
	"github.com/uchupx/kajian-auth/internal/handler"
	"github.com/uchupx/kajian-auth/internal/repo"
	"github.com/uchupx/kajian-auth/internal/service"
	"github.com/uchupx/kajian-auth/internal/service/jwt"
)

type Internal struct {
	// adapter
	db          *sqlx.DB
	redisClient *redis.Client

	// repo
	userRepo *repo.UserRepo

	// handler
	authHandler *handler.AuthHandler

	// service
	userService *service.UserService
	jwtService  jwt.CryptService
}

func (i *Internal) DB(conf *config.Config) *sqlx.DB {
	if i.db == nil {
		db, err := mysql.NewConnection(mysql.DBPayload{
			Host:     conf.Database.Host,
			Port:     conf.Database.Port,
			Username: conf.Database.User,
			Password: conf.Database.Password,
			Database: conf.Database.Name,
		})
		if err != nil {
			panic(err)
		}

		i.db = db
	}
	return i.db
}

func (i *Internal) RedisClient(conf *config.Config) *redis.Client {
	if i.redisClient == nil {
		redisClient, err := kajianRedis.NewRedisConn(kajianRedis.RedisConfig{
			Host:     conf.Redis.Host,
			Password: conf.Redis.Password,
			Database: 0,
		})
		if err != nil {
			panic(err)
		}

		i.redisClient = redisClient

	}

	return i.redisClient
}

func (i *Internal) AuthHandler(conf *config.Config) *handler.AuthHandler {
	if i.authHandler == nil {
		i.authHandler = &handler.AuthHandler{
			UserService: i.UserService(conf),
		}
	}

	return i.authHandler
}

func (i *Internal) UserService(conf *config.Config) *service.UserService {
	if i.userService == nil {
		i.userService = &service.UserService{
			UserRepo: i.UserRepo(conf),
			JWT:      i.JWTService(conf),
		}
	}

	return i.userService
}

func (i *Internal) JWTService(conf *config.Config) jwt.CryptService {
	if i.jwtService == nil {
		rsa := jwt.NewCryptService(jwt.Params{
			Conf: conf,
		})

		i.jwtService = rsa
	}

	return i.jwtService
}

func (i *Internal) UserRepo(conf *config.Config) *repo.UserRepo {
	if i.userRepo == nil {
		i.userRepo = repo.NewUserRepo(i.DB(conf))
	}

	return i.userRepo
}

func (i *Internal) InitRoutes(conf *config.Config, e *echo.Echo) {
	routes := []handler.BaseHandler{i.AuthHandler(conf)}

	for _, route := range routes {
		route.InitRoutes(e)
	}
}
