package elasticsearch

import (
	"context"
	"fmt"
	"strconv"

	elastic "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/pdc/models"
	"github.com/lib/pq"
)

//VideoIndex 视频索引
type VideoIndex struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Desc        string         `json:"desc"`
	PubDate     int64          `json:"pub_date"`
	Cover       string         `json:"cover"`
	TotalNum    int32          `json:"total_num"`
	Tags        pq.StringArray `json:"tags"`
	IsShow      bool           `json:"is_show"`
	SeriesID    uint           `json:"series_id"`
	SeriesName  string         `json:"series_name"`
	SeriesAlias string         `json:"series_alias"`
	SeriesNum   uint           `json:"series_num"`
}

//GetByID ..
func (v *VideoIndex) GetByID(id string) error {
	return models.Gorm.Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,b.total_num,a.tags,a.is_show,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(models.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+models.TablePrefix+"_episode where video_id=? group by video_id) b on a.id=b.video_id", id).
		Joins("left join "+models.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+models.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Where("a.id=?", id).Take(v).Error
}

//Find ..
func (v *VideoIndex) Find() ([]*VideoIndex, error) {
	data := make([]*VideoIndex, 0)
	err := models.Gorm.Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,b.total_num,a.tags,a.is_show,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(models.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+models.TablePrefix+"_episode  group by video_id) b on a.id=b.video_id").
		Joins("left join "+models.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+models.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Where("a.is_show=?", true).
		Find(&data).Error
	return data, err
}

//BulkSaveES 批量保存到es
func (v *VideoIndex) BulkSaveES(ctx context.Context, vis []*VideoIndex, indexName string, bulkNum, workerNum int) {
	bds := make([]*elastic.BulkDoc, 0, len(vis))
	for _, v := range vis {
		bd := &elastic.BulkDoc{
			ID:  strconv.Itoa(int(v.ID)),
			Doc: v,
		}
		bds = append(bds, bd)
	}
	errs := ESClient.BulkInsert(ctx, bds, indexName, bulkNum, workerNum)
	for _, v := range errs {
		fmt.Println(v)
	}
}

//VideoMapping ..
const VideoMapping = `{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0,
        "analysis": {
            "analyzer": {
                "ik_max_synonym": {
                    "type": "custom",
                    "tokenizer": "ik_max_word",
                    "filter": [
                        "my_filter"
                    ]
                },
                "ik_smart_synonym": {
                    "type": "custom",
                    "tokenizer": "ik_smart",
                    "filter": [
                        "my_filter"
                    ]
                }
            },
            "filter": {
                "my_filter": {
                    "type": "synonym",
                    "synonyms_path": "analysis/synonym.txt"
                }
            }
        }
    },
    "mappings": {
        "include_in_all": "false",
        "dynamic": true,
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "desc": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 5000
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "pub_date": {
                "type": "long"
            },
            "cover": {
                "type": "keyword"
            },
            "total_num": {
                "type": "long"
            },
            "tags": {
                "type": "keyword"
            },
            "series_id": {
                "type": "long"
            },
            "series_name": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "series_alias": {
                "type": "text",
                "fields": {
                    "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                    }
                },
                "analyzer": "ik_max_synonym",
                "search_analyzer": "ik_smart_synonym"
            },
            "series_num": {
                "type": "long"
            }
        }
    }
}`
