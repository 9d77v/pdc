package consumers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/9d77v/go-pkg/db/elastic"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/module/video-service/models"

	"github.com/nats-io/nats.go"
)

//HandleVideoMSG ...
func HandleVideoMSG(m *nats.Msg) {
	ctx := context.Background()
	id := string(m.Data)
	client := elastic.GetClient()
	vi := new(models.VideoIndex)
	indexNames := client.FindIndexesByAlias(ctx,
		consts.AliasVideo, consts.ESLayout)
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

func syncAllVideos(ctx context.Context, vi *models.VideoIndex, client *elastic.Client) error {
	indexName := client.GetNewIndexName(consts.AliasVideo, consts.ESLayout)
	err := client.CreateIndex(ctx, indexName, consts.VideoMapping)
	if err != nil {
		return err
	}
	data, err := vi.Find()
	if err != nil {
		return err
	}
	bulkSaveES(ctx, data, indexName, 1000, 3)
	err = client.SetNewAlias(ctx, consts.AliasVideo, indexName)
	if err != nil {
		return err
	}
	indexNames := client.FindIndexesByAlias(ctx, consts.AliasVideo, consts.ESLayout)
	return client.KeepIndex(ctx, indexNames, 3)
}

func bulkSaveES(ctx context.Context,
	vis []*models.VideoIndex, indexName string, bulkNum, workerNum int) {
	bds := make([]*elastic.BulkDoc, 0, len(vis))
	for _, v := range vis {
		bd := &elastic.BulkDoc{
			ID:  strconv.Itoa(int(v.ID)),
			Doc: v,
		}
		bds = append(bds, bd)
	}
	errs := elastic.GetClient().BulkInsert(ctx, bds, indexName, bulkNum, workerNum)
	for _, v := range errs {
		fmt.Println(v)
	}
}

func syncOneVideoRecord(ctx context.Context, id string, vi *models.VideoIndex,
	client *elastic.Client) error {
	videoID, _ := strconv.ParseUint(id, 10, 64)
	err := vi.GetByID(uint(videoID))
	if err != nil {
		return err
	}
	_, err = client.Index().
		Index(consts.AliasVideo).
		Id(id).
		BodyJson(vi).
		Do(ctx)
	return err
}
