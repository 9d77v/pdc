package services

import (
	"log"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/models"
)

//GetPageInfo 获取分页信息
func GetPageInfo(page, pageSize *int64) (int, int) {
	offset := ptrs.Int64(page)
	limit := ptrs.Int64(pageSize)
	if offset < 1 {
		offset = 1
	}
	if limit == 0 {
		limit = 10
	}
	if limit < 0 {
		limit = -1
	}
	if limit > 100 {
		limit = 100
	}
	offset = (offset - 1) * limit
	return int(offset), int(limit)
}

//GetElasticPageInfo 获取Elasticsearch分页信息
func GetElasticPageInfo(page, pageSize *int64) (int, int) {
	offset := ptrs.Int64(page)
	limit := ptrs.Int64(pageSize)
	if offset < 1 {
		offset = 1
	}
	if limit <= 0 {
		limit = 500
	}
	if limit > 500 {
		limit = 500
	}
	offset = (offset - 1) * limit
	return int(offset), int(limit)
}

//AckHandler ..
func AckHandler(ackedNuid string, err error) {
	if err != nil {
		log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
	} else if models.DEBUG {
		log.Printf("Received ack for msg id %s\n", ackedNuid)
	}
}
