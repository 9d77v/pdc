package consumers

import (
	"context"
	"log"

	v7 "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/pdc/internal/db/elasticsearch"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/9d77v/pdc/internal/module/video-service/services"

	"github.com/nats-io/stan.go"
)

var videoSearch = new(services.VideoSearch)

//HandleVideoMSG ...
func HandleVideoMSG(m *stan.Msg) {
	ctx := context.Background()
	vi := new(models.VideoIndex)
	client := elasticsearch.GetClient()

	indexNames := client.FindIndexesByAlias(ctx,
		elasticsearch.AliasVideo, elasticsearch.ESLayout)
	id := string(m.Data)
	if string(m.Data) == "0" || len(indexNames) == 0 {
		err := syncAllVideos(ctx, vi, client)
		if err != nil {
			log.Println("syncOneVideoRecord  error:", err)
		}
		return
	}
	err := syncOneVideoRecord(ctx, id, vi, client)
	if err != nil {
		log.Println("syncOneVideoRecord  error:", err)
	}
}

func syncAllVideos(ctx context.Context, vi *models.VideoIndex, client *v7.Client) error {
	indexName := client.GetNewIndexName(elasticsearch.AliasVideo, elasticsearch.ESLayout)
	err := client.CreateIndex(ctx, indexName, elasticsearch.VideoMapping)
	if err != nil {
		return err
	}
	data, err := vi.Find()
	if err != nil {
		return err
	}
	videoSearch.BulkSaveES(ctx, data, indexName, 1000, 3)
	err = client.SetNewAlias(ctx, elasticsearch.AliasVideo, indexName)
	if err != nil {
		return err
	}
	indexNames := client.FindIndexesByAlias(ctx, elasticsearch.AliasVideo, elasticsearch.ESLayout)
	return client.KeepIndex(ctx, indexNames, 3)
}

func syncOneVideoRecord(ctx context.Context, id string, vi *models.VideoIndex,
	client *v7.Client) error {
	err := vi.GetByID(id)
	if err != nil {
		return err
	}
	_, err = client.Index().
		Index(elasticsearch.AliasVideo).
		Id(id).
		BodyJson(vi).
		Do(ctx)
	return err
}
