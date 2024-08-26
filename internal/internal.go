package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-api/pkg/db"
	"github.com/uchupx/kajian-api/pkg/logger"
	"github.com/uchupx/kajian-api/pkg/mysql"
	kajianRedis "github.com/uchupx/kajian-api/pkg/redis"
	"github.com/uchupx/kajian-auth/config"
	"github.com/uchupx/kajian-auth/internal/handler"
	"github.com/uchupx/kajian-auth/internal/middlware"
	"github.com/uchupx/kajian-auth/internal/repo"
	"github.com/uchupx/kajian-auth/internal/service"
	"github.com/uchupx/kajian-auth/internal/service/jwt"
	"github.com/uchupx/kajian-auth/pb"
	"google.golang.org/grpc"
)

type Internal struct {
	// adapter
	db          *db.DB
	redisClient *redis.Client

	// repo
	userRepo   *repo.UserRepo
	clientRepo *repo.ClientRepo

	//middlware
	middlware *middlware.Middleware

	// handler
	authHandler     *handler.AuthHandler
	authGRPCHandler *handler.AuthGRPCHandler

	// service
	userService *service.UserService
	jwtService  jwt.CryptService
}

func (i *Internal) DB(conf *config.Config) *db.DB {
	if i.db == nil {
		fmt.Printf("test")
		for idx := 1; idx <= 3; idx++ {
			db, err := mysql.NewConnection(mysql.DBPayload{
				Host:     conf.Database.Host,
				Port:     conf.Database.Port,
				Username: conf.Database.User,
				Password: conf.Database.Password,
				Database: conf.Database.Name,
			})
			if err != nil {
				fmt.Println(err)
				time.Sleep(2 * time.Second)
				fmt.Println("Retry connection")

				if idx == 3 {
					panic(err)
				}

				continue
			}

			db.SetDebug(i.isDebug(conf))
			i.db = db
			break
		}
	}
	return i.db
}

func (i *Internal) RedisClient(conf *config.Config) *redis.Client {
	if i.redisClient == nil {
		redisClient, err := kajianRedis.NewRedisConn(kajianRedis.RedisConfig{
			Host:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
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

func (i *Internal) AuthGRPCHandler(conf *config.Config) *handler.AuthGRPCHandler {
	if i.authGRPCHandler == nil {
		i.authGRPCHandler = &handler.AuthGRPCHandler{
			UserService: i.UserService(conf),
		}
	}

	return i.authGRPCHandler
}

func (i *Internal) UserService(conf *config.Config) *service.UserService {
	if i.userService == nil {
		i.userService = &service.UserService{
			UserRepo:   i.UserRepo(conf),
			JWT:        i.JWTService(conf),
			Redis:      i.RedisClient(conf),
			ClientRepo: i.ClientRepo(conf),
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

func (i *Internal) ClientRepo(conf *config.Config) *repo.ClientRepo {
	if i.clientRepo == nil {
		i.clientRepo = repo.NewClientRepo(i.DB(conf))
	}

	return i.clientRepo
}

func (i *Internal) Middlware(conf *config.Config) *middlware.Middleware {
	if i.middlware == nil {
		i.middlware = middlware.New(middlware.Config{})
	}

	return i.middlware
}

func (i *Internal) InitRoutes(conf *config.Config, e *echo.Echo) {
	routes := []handler.BaseHandler{i.AuthHandler(conf)}
	i.Middlware(conf)
	// e.Use(i.middlware.Logger)
	// eMiddleware.CORSWithConfig(eMiddleware.CORSConfig{})
	e.Use(i.middlware.Recover)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	for _, route := range routes {
		route.InitRoutes(e)
	}
}

func (i *Internal) InitRoutesGRPC(conf *config.Config, s *grpc.Server) {
	pb.RegisterAuthorizationServiceServer(s, i.AuthGRPCHandler(conf))
}

func (i *Internal) InitLogger(conf *config.Config) {
	logConf := logger.LogConfig{
		Path:    conf.App.Log,
		NameApp: "kajian-auth",
	}

	logConf.SetDebug(i.isDebug(conf))
	logger.InitLog(logConf)
}

func (i *Internal) isDebug(conf *config.Config) bool {
	return conf.App.Env == "dev"
}
