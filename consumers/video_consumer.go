package consumers

import (
	"context"
	"log"

	"github.com/9d77v/pdc/models/es"
	"github.com/nats-io/stan.go"
)

//HandleVideoMSG ...
func HandleVideoMSG(m *stan.Msg) {
	ctx := context.Background()
	client := es.ESClient
	indexNames := client.FindIndexesByAlias(ctx,
		es.AliasVideo, es.ESLayout)
	id := string(m.Data)
	vi := new(es.VideoIndex)
	if string(m.Data) == "0" || len(indexNames) == 0 {
		indexName := client.GetNewIndexName(es.AliasVideo, es.ESLayout)
		err := client.CreateIndex(ctx, indexName, es.VideoMapping)
		if err != nil {
			log.Println("create index error:", err)
			return
		}
		data, err := vi.Find()
		if err != nil {
			log.Println("get data error:", err)
			return
		}
		vi.BulkSaveES(ctx, data, indexName, 1000, 3)
		client.SetNewAlias(ctx, es.AliasVideo, indexName)
		indexNames = client.FindIndexesByAlias(ctx, es.AliasVideo, es.ESLayout)
		client.KeepIndex(ctx, indexNames, 3)
	} else {
		err := vi.GetByID(id)
		if err != nil {
			log.Println("get data error:", err)
			return
		}
		_, err = client.Index().
			Index(es.AliasVideo).
			Id(id).
			BodyJson(vi).
			Do(ctx)
		if err != nil {
			log.Println("insert data error:", err)
			return
		}
	}
}
