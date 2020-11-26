package redis

import (
	"sync"

	"github.com/9d77v/pdc/internal/utils"
	redisGo "github.com/go-redis/redis/v8"
)

//环境变量
var (
	redisAddress  = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	client *redisGo.Client
	once   sync.Once
)

//redis前缀
const (
	PrefixUser = "USER"
)

//GetClient get redis connection
func GetClient() *redisGo.Client {
	once.Do(func() {
		client = initClient()
	})
	return client
}
func initClient() *redisGo.Client {
	return redisGo.NewClient(&redisGo.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}
