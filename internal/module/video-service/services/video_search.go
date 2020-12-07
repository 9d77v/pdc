package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	elastic "github.com/olivere/elastic/v7"

	es "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/elasticsearch"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/9d77v/pdc/internal/utils"
)

//VideoSearch ..
type VideoSearch struct {
}

//BulkSaveES 批量保存到es
func (s VideoSearch) BulkSaveES(ctx context.Context,
	vis []*models.VideoIndex, indexName string, bulkNum, workerNum int) {
	bds := make([]*es.BulkDoc, 0, len(vis))
	for _, v := range vis {
		bd := &es.BulkDoc{
			ID:  strconv.Itoa(int(v.ID)),
			Doc: v,
		}
		bds = append(bds, bd)
	}
	errs := elasticsearch.GetClient().BulkInsert(ctx, bds, indexName, bulkNum, workerNum)
	for _, v := range errs {
		fmt.Println(v)
	}
}

//GetByIDFromElastic ...
func (s VideoSearch) GetByIDFromElastic(ctx context.Context,
	videoID string) (*models.VideoIndex, error) {
	result, err := elasticsearch.GetClient().Get().Index(elasticsearch.AliasVideo).Id(videoID).Do(ctx)
	if err != nil {
		return nil, err
	}
	vi := new(models.VideoIndex)
	data, err := result.Source.MarshalJSON()
	if err != nil {
		log.Println("elastic search result json marshal error:", err)
		return nil, err
	}
	err = json.Unmarshal(data, &vi)
	if err != nil {
		log.Println("elastic search result json unmarshal error:", err)
		return nil, err
	}
	return vi, nil
}

//ListVideoIndex ..
func (s VideoSearch) ListVideoIndex(ctx context.Context,
	searchParam model.SearchParam, scheme string) (int64, []*model.VideoIndex, []*model.AggResult, error) {
	boolQuery := elastic.NewBoolQuery()
	keywordStr := strings.ReplaceAll(ptrs.String(searchParam.Keyword), " ", "")
	if keywordStr != "" {
		s.buildKeywordQuery(keywordStr, boolQuery)
	}
	boolQuery.Filter(s.buildFilterQuery(searchParam))
	offset, limit := utils.GetElasticPageInfo(searchParam.Page, searchParam.PageSize)
	searchService := elasticsearch.GetClient().Search().
		Index(elasticsearch.AliasVideo).
		Query(boolQuery).
		From(int(offset)).
		Size(int(limit))
	aggsParams := []*es.AggsParam{
		{Field: "tags", Size: 50},
	}
	field := base.NewGraphQLField(ctx, "")
	if field.FieldMap["aggResults"] {
		searchService = es.Aggs(searchService, aggsParams...)
	}
	if field.FieldMap["edges"] {
		if ptrs.Bool(searchParam.IsRandom) {
			searchService.SortBy(elastic.NewScriptSort(elastic.NewScript("Math.random()"), "number").Order(true))
		} else {
			searchService.Sort("_score", false).
				Sort("title.keyword", true)
		}
	} else {
		searchService.FetchSource(false)
	}
	result, err := searchService.Do(ctx)
	if err != nil {
		log.Println("err:", err)
		return 0, []*model.VideoIndex{}, []*model.AggResult{}, nil
	}
	return result.TotalHits(), s.getData(field.FieldMap, result, scheme),
		s.getAggResults(field, aggsParams, result), nil
}

func (s VideoSearch) buildKeywordQuery(keywordStr string, boolQuery *elastic.BoolQuery) {
	subBoolQuery := elastic.NewMultiMatchQuery(keywordStr, []string{
		"title^10",
		"series_name^5",
		"series_alias^0.5",
		"desc^0.1",
		"title.ikmax^10",
		"series_name.ikmax^5",
		"series_alias.ikmax^0.5",
		"title.sy_ikmax^10",
		"series_name.sy_ikmax^5",
		"series_alias.sy_ikmax^0.5",
		"title.synonym^10",
		"series_name.synonym^5",
		"series_alias.synonym^0.5",
	}...).
		Type("cross_fields").
		Operator("AND").
		TieBreaker(0.3)
	boolQuery.Must(subBoolQuery)
}

func (s VideoSearch) buildFilterQuery(searchParam model.SearchParam,
	etcIgnoreQueries ...elastic.Query) *elastic.BoolQuery {
	filterQueries := make([]elastic.Query, 0)
	filterQueries = append(filterQueries, elastic.NewTermQuery("is_show", true))
	if len(searchParam.Tags) > 0 {
		for _, v := range searchParam.Tags {
			if v != "" {
				filterQueries = append(filterQueries, elastic.NewTermQuery("tags", v))
			}
		}
	}
	ignoreQueries := make([]elastic.Query, 0)
	if ptrs.Bool(searchParam.IsMobile) {
		ignoreQueries = append(ignoreQueries, elastic.NewTermQuery("is_hide_on_mobile", true))
	}
	return elastic.NewBoolQuery().
		Must(filterQueries...).
		MustNot(append(ignoreQueries, etcIgnoreQueries...)...)
}

func (s VideoSearch) getData(fieldMap map[string]bool, result *elastic.SearchResult,
	scheme string) []*model.VideoIndex {
	vis := make([]*model.VideoIndex, 0)
	if fieldMap["edges"] {
		for _, v := range result.Hits.Hits {
			vi := new(models.VideoIndex)
			data, err := v.Source.MarshalJSON()
			if err != nil {
				log.Println("elastic search result json marshal error:", err)
			}
			err = json.Unmarshal(data, &vi)
			if err != nil {
				log.Println("elastic search result json unmarshal error:", err)
			}
			vi.Cover = oss.GetOSSPrefixByScheme(scheme) + vi.Cover
			vis = append(vis, s.getVideoIndex(vi))
		}
	}
	return vis
}

func (s VideoSearch) getAggResults(field base.GraphQLField, aggsParams []*es.AggsParam,
	result *elastic.SearchResult) []*model.AggResult {
	aggResults := make([]*model.AggResult, 0)
	if field.FieldMap["aggResults"] {
		for _, v := range aggsParams {
			aggResult, found := result.Aggregations.Terms("group_by_" + v.Field)
			if found {
				for _, v := range aggResult.Buckets {
					aggResults = append(aggResults, &model.AggResult{
						Key:   v.Key.(string),
						Value: v.DocCount,
					})
				}
			}
		}
	}
	return aggResults
}

//SimilarVideoIndex ..
func (s VideoSearch) SimilarVideoIndex(ctx context.Context,
	searchParam model.SearchParam, episodeID int64, scheme string) (int64, []*model.VideoIndex, error) {
	videoID := models.NewEpisode().GetVideoIDByID(uint(episodeID))
	id := strconv.FormatUint(uint64(videoID), 10)
	videoDoc := new(models.VideoIndex)
	err := videoDoc.GetByID(id)
	if err != nil {
		return 0, []*model.VideoIndex{}, nil
	}
	seriesID := strconv.FormatUint(uint64(videoDoc.SeriesID), 10)
	boolQuery := elastic.NewBoolQuery().
		Must(s.buildFunctionScoreQuery(id)).
		Filter(s.buildFilterQuery(searchParam, elastic.NewTermQuery("series_id", seriesID)))
	searchService := elasticsearch.GetClient().Search().
		Index(elasticsearch.AliasVideo).
		Query(boolQuery).
		Size(int(ptrs.Int64(searchParam.PageSize))).
		Sort("_score", false).
		Sort("title.keyword", true)
	result, err := searchService.Do(ctx)
	if err != nil {
		log.Println("err:", err)
		return 0, []*model.VideoIndex{}, nil
	}
	return result.TotalHits(), s.getData(map[string]bool{"edges": true}, result, scheme), nil
}

func (s VideoSearch) buildFunctionScoreQuery(videoID string) *elastic.FunctionScoreQuery {
	fsQuery := elastic.NewFunctionScoreQuery()
	moreLikeThisItem := elastic.NewMoreLikeThisQueryItem().Index(elasticsearch.AliasVideo).Id(videoID)
	mltQuery := elastic.NewMoreLikeThisQuery().
		Field("title", "desc", "tags").
		LikeItems(moreLikeThisItem).
		MinTermFreq(1).
		MinDocFreq(5).
		StopWord("的", "第一季", "第二季", "第三季").
		MaxQueryTerms(12).
		Analyzer("ik_smart_synonym")
	return fsQuery.Query(mltQuery)
}
