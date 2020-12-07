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
	input model.SearchParam, scheme string) (int64, []*model.VideoIndex, []*model.AggResult, error) {
	boolQuery := elastic.NewBoolQuery()
	keywordStr := strings.ReplaceAll(ptrs.String(input.Keyword), " ", "")
	if keywordStr != "" {
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
	offset, limit := utils.GetElasticPageInfo(input.Page, input.PageSize)
	filterQueries := make([]elastic.Query, 0)
	filterQueries = append(filterQueries, elastic.NewTermQuery("is_show", true))

	if len(input.Tags) > 0 {
		for _, v := range input.Tags {
			if v != "" {
				filterQueries = append(filterQueries, elastic.NewTermQuery("tags", v))
			}
		}
	}

	ignoreQueries := make([]elastic.Query, 0)
	if ptrs.Bool(input.IsMobile) {
		ignoreQueries = append(ignoreQueries, elastic.NewTermQuery("is_hide_on_mobile", true))
	}
	filterQuery := elastic.NewBoolQuery().
		Must(filterQueries...).MustNot(ignoreQueries...)
	boolQuery.Filter(filterQuery)
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
		if ptrs.Bool(input.IsRandom) {
			searchService.SortBy(elastic.NewScriptSort(elastic.NewScript("Math.random()"), "number").Order(true))
		} else {
			searchService.Sort("_score", false).
				Sort("title.keyword", true)
		}
	} else {
		searchService.FetchSource(false)
	}
	result, err := searchService.Do(ctx)
	vis := make([]*model.VideoIndex, 0)
	aggResults := make([]*model.AggResult, 0)
	if err != nil {
		log.Println("err:", err)
		return 0, vis, aggResults, nil
	}
	if field.FieldMap["edges"] {
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
	return result.TotalHits(), vis, aggResults, nil
}

//SimilarVideoIndex ..
func (s VideoSearch) SimilarVideoIndex(ctx context.Context,
	searchParam model.SearchParam, episodeID int64, scheme string) (int64, []*model.VideoIndex, error) {
	videoID := models.NewEpisode().GetVideoIDByID(uint(episodeID))
	id := strconv.FormatUint(uint64(videoID), 10)
	vis := make([]*model.VideoIndex, 0)

	boolQuery := elastic.NewBoolQuery()
	fsQuery := elastic.NewFunctionScoreQuery()
	mltQuery := elastic.NewMoreLikeThisQuery()
	moreLikeThisItem := elastic.NewMoreLikeThisQueryItem()
	moreLikeThisItem = moreLikeThisItem.Index(elasticsearch.AliasVideo).Id(id)
	stopWords := []string{"的", "第一季", "第二季", "第三季"}
	mltQuery = mltQuery.Field(
		"title",
		"desc",
		"tags").
		LikeItems(moreLikeThisItem).
		MinTermFreq(1).
		MinDocFreq(5).StopWord(
		stopWords...,
	).
		MaxQueryTerms(12).
		Analyzer("ik_smart_synonym")
	videoDoc := new(models.VideoIndex)
	err := videoDoc.GetByID(id)
	if err != nil {
		return 0, vis, nil
	}
	filterQueries := make([]elastic.Query, 0)
	filterQueries = append(filterQueries, elastic.NewTermQuery("is_show", true))

	ignoreQueries := make([]elastic.Query, 0)
	seriesID := strconv.FormatUint(uint64(videoDoc.SeriesID), 10)
	ignoreQueries = append(ignoreQueries, elastic.NewTermQuery("series_id", seriesID))
	if ptrs.Bool(searchParam.IsMobile) {
		ignoreQueries = append(ignoreQueries, elastic.NewTermQuery("is_hide_on_mobile", true))
	}
	filterQuery := elastic.NewBoolQuery().
		Must(filterQueries...).
		MustNot(ignoreQueries...)
	boolQuery.Must(fsQuery.Query(mltQuery)).Filter(filterQuery)
	searchService := elasticsearch.GetClient().Search().
		Index(elasticsearch.AliasVideo).
		Query(boolQuery).
		Size(int(ptrs.Int64(searchParam.PageSize))).
		Sort("_score", false).
		Sort("title.keyword", true)
	result, err := searchService.Do(ctx)
	if err != nil {
		log.Println("err:", err)
		return 0, vis, nil
	}
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
	return result.TotalHits(), vis, nil
}
