package consumer

import (
	"context"
	"log"

	"github.com/9d77v/pdc/internal/db/elasticsearch"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/9d77v/pdc/internal/module/video-service/services"

	"github.com/nats-io/stan.go"
)

var videoSearch = new(services.VideoSearch)

//HandleVideoMSG ...
func HandleVideoMSG(m *stan.Msg) {
	ctx := context.Background()
	client := elasticsearch.ESClient
	indexNames := client.FindIndexesByAlias(ctx,
		elasticsearch.AliasVideo, elasticsearch.ESLayout)
	id := string(m.Data)
	vi := new(models.VideoIndex)
	if string(m.Data) == "0" || len(indexNames) == 0 {
		indexName := client.GetNewIndexName(elasticsearch.AliasVideo, elasticsearch.ESLayout)
		err := client.CreateIndex(ctx, indexName, elasticsearch.VideoMapping)
		if err != nil {
			log.Println("create index error:", err)
			return
		}
		data, err := vi.Find()
		if err != nil {
			log.Println("get data error:", err)
			return
		}
		videoSearch.BulkSaveES(ctx, data, indexName, 1000, 3)
		err = client.SetNewAlias(ctx, elasticsearch.AliasVideo, indexName)
		if err != nil {
			log.Println("SetNewAlias  error:", err)
			return
		}
		indexNames = client.FindIndexesByAlias(ctx, elasticsearch.AliasVideo, elasticsearch.ESLayout)
		err = client.KeepIndex(ctx, indexNames, 3)
		if err != nil {
			log.Println("KeepIndex  error:", err)
			return
		}
	} else {
		err := vi.GetByID(id)
		if err != nil {
			log.Println("get data error:", err)
			return
		}
		_, err = client.Index().
			Index(elasticsearch.AliasVideo).
			Id(id).
			BodyJson(vi).
			Do(ctx)
		if err != nil {
			log.Println("insert data error:", err)
			return
		}
	}
}
