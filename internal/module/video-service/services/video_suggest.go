package services

import (
	"context"
	"log"
	"strconv"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/elasticsearch"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	elastic "github.com/olivere/elastic/v7"
)

type videoSuggest struct {
	*base.Search
	episodeID int64
}

func newVideoSuggest(ctx context.Context,
	searchParam model.SearchParam, episodeID int64, scheme string) *videoSuggest {
	return &videoSuggest{
		Search:    base.NewSearch(ctx, searchParam, scheme),
		episodeID: episodeID,
	}
}

func (s *videoSuggest) execute() (*model.VideoIndexConnection, error) {
	result := new(model.VideoIndexConnection)
	videoID := models.NewEpisode().GetVideoIDByID(uint(s.episodeID))
	s.BoolQuery.Must(newFunctionScoreQuery(videoID))
	videoDoc := new(models.VideoIndex)
	err := videoDoc.GetByID(videoID)
	if err != nil {
		return result, err
	}
	if videoDoc.SeriesID != 0 {
		seriesID := strconv.FormatUint(uint64(videoDoc.SeriesID), 10)
		s.IgnoreQueries = append(s.IgnoreQueries, elastic.NewTermQuery("series_id", seriesID))
	}
	s.Filter()
	searchService := elasticsearch.GetClient().Search().
		Index(elasticsearch.AliasVideo).
		Query(s.BoolQuery).
		Size(int(ptrs.Int64(s.SearchParam.PageSize))).
		Sort("_score", false).
		Sort("title.keyword", true)
	searchResult, err := searchService.Do(s.Ctx)
	if err != nil {
		log.Println("err:", err)
		return result, err
	}
	result.TotalCount = searchResult.TotalHits()
	result.Edges = getEdges(searchResult, s.Scheme)
	return result, nil
}

func newFunctionScoreQuery(videoID uint) *elastic.FunctionScoreQuery {
	fsQuery := elastic.NewFunctionScoreQuery()
	id := strconv.FormatUint(uint64(videoID), 10)
	moreLikeThisItem := elastic.NewMoreLikeThisQueryItem().Index(elasticsearch.AliasVideo).Id(id)
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
