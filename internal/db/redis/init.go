package redis

import (
	"github.com/9d77v/pdc/internal/utils"
	redisGo "github.com/go-redis/redis/v8"
)

//环境变量
var (
	redisAddress  = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	//Client ..
	Client *redisGo.Client
)

//redis前缀
const (
	PrefixUser = "USER"
)

func init() {
	Client = redisGo.NewClient(&redisGo.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}
