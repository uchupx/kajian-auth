package middlware

import "github.com/redis/go-redis/v9"

type Middleware struct {
	Redis *redis.Client
}

type Config struct {
	Redis *redis.Client
}

func New(conf Config) *Middleware {
	return &Middleware{
		Redis: conf.Redis,
	}
}

type logMeta struct {
	Methods  string `json:"methods"`
	Path     string `json:"path"`
	Latency  string `json:"latency"`
	Status   int    `json:"status"`
	Response string `json:"response"`
}
