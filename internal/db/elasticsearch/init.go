package elasticsearch

import (
	"strings"
	"sync"

	"github.com/9d77v/go-lib/clients/config"
	elastic "github.com/9d77v/go-lib/clients/elastic/v7"

	"github.com/9d77v/pdc/internal/utils"
)

var (
	client  *elastic.Client
	esAddrs = utils.GetEnvStr("ESADDR", "http://domain.local:9200")
	once    sync.Once
)

//别名
const (
	AliasVideo = "video"
	ESLayout   = "2006.01.02-15.04.05"
)

//GetClient get elasticsearch connection
func GetClient() *elastic.Client {
	once.Do(func() {
		client = initClient()
	})
	return client
}

func initClient() *elastic.Client {
	db, err := elastic.NewClient(&config.ElasticConfig{URLs: strings.Split(esAddrs, ",")})
	if err != nil {
		// Handle error
		panic(err)
	}
	return db
}
