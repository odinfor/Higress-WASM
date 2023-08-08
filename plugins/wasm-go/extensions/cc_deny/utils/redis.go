package utils

/*单机redis*/

import (
	"cc_deny/types"
	"github.com/redis/go-redis/v9"
)

type (
	Redis struct {
		addr     string
		password string
		db       int
		client   *redis.Client
	}
)

func NewRedis() *Redis {
	rConf := types.NewConfDo().RedisConf()
	return &Redis{
		addr:     rConf.Addr,
		password: rConf.Password,
		db:       rConf.DB,
		client: redis.NewClient(&redis.Options{
			Addr:     rConf.Addr,
			Password: rConf.Password,
			DB:       rConf.DB,
		}),
	}
}

func (r *Redis) Client() *redis.Client {
	return r.client
}
