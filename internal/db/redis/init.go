package redis

import (
	"strings"
	"sync"

	"github.com/9d77v/pdc/internal/utils"
	redisGo "github.com/go-redis/redis/v8"
)

//环境变量
var (
	redisAddresses = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword  = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	client redisGo.Cmdable
	once   sync.Once
)

//redis前缀
const (
	PrefixUser              = "USER"
	PrefixVideoDataUser     = "VIDEO_DATA:USER"
	PrefixVideoDataAnime    = "VIDEO_DATA:ANIME"
	PrefixVideoDataEpisode  = "VIDEO_DATA:EPISODE"
	PrefixVideoDataDuration = "VIDEO_DATA:DURATION"
)

//GetClient get redis connection
func GetClient() redisGo.Cmdable {
	once.Do(func() {
		client = initClient()
	})
	return client
}

func initClient() redisGo.Cmdable {
	addresses := strings.Split(redisAddresses, ";")
	if len(addresses) == 1 {
		return redisGo.NewClient(&redisGo.Options{
			Addr:     addresses[0],
			Password: redisPassword,
		})
	}
	return redisGo.NewClusterClient(&redisGo.ClusterOptions{
		Addrs:    addresses,
		Password: redisPassword,
	})
}
