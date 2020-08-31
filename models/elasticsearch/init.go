package elasticsearch

import (
	"github.com/9d77v/go-lib/clients/config"
	elastic "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/pdc/utils"
)

var (
	//ESClient Elasticsearch Client
	ESClient *elastic.Client
	//ESAddr Elasticsearch Server Address
	ESAddr = utils.GetEnvStr("ESADDR", "http://domain.local:9200")
)

//别名
const (
	AliasVideo = "video"
	ESLayout   = "2006.01.02-15.04.05"
)

func init() {
	var err error
	ESClient, err = elastic.NewClient(&config.ElasticConfig{URLs: []string{ESAddr}})
	if err != nil {
		// Handle error
		panic(err)
	}
}
