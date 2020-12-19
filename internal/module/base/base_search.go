package base

import (
	"context"

	es "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/pdc/internal/graph/model"
	elastic "github.com/olivere/elastic/v7"
)

//Search ..
type Search struct {
	Ctx           context.Context
	SearchParam   *SearchParam
	Scheme        string
	SearchService *elastic.SearchService
	BoolQuery     *elastic.BoolQuery
	FilterQueries []elastic.Query
	IgnoreQueries []elastic.Query
}

//NewSearch ..
func NewSearch(ctx context.Context, searchParam *SearchParam, scheme string) *Search {
	return &Search{
		Ctx:           ctx,
		SearchParam:   searchParam,
		Scheme:        scheme,
		BoolQuery:     elastic.NewBoolQuery(),
		FilterQueries: newFilterQueries(searchParam),
		IgnoreQueries: newIgnoreQueries(searchParam),
	}
}

func newFilterQueries(searchParam *SearchParam) []elastic.Query {
	mustQueries := make([]elastic.Query, 0)
	mustQueries = append(mustQueries, elastic.NewTermQuery("is_show", true))
	if len(searchParam.Tags) > 0 {
		for _, v := range searchParam.Tags {
			if v != "" {
				mustQueries = append(mustQueries, elastic.NewTermQuery("tags", v))
			}
		}
	}
	return mustQueries
}

func newIgnoreQueries(searchParam *SearchParam) []elastic.Query {
	mustNotQueries := make([]elastic.Query, 0)
	if searchParam.IsMobile {
		mustNotQueries = append(mustNotQueries, elastic.NewTermQuery("is_hide_on_mobile", true))
	}
	return mustNotQueries
}

//Filter ..
func (s *Search) Filter() {
	s.BoolQuery.Filter(elastic.NewBoolQuery().
		Must(s.FilterQueries...).
		MustNot(s.IgnoreQueries...))
}

//GetPageInfo ..
func (s *Search) GetPageInfo() (int, int) {
	page := s.SearchParam.Page
	pageSize := s.SearchParam.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 500 {
		pageSize = 500
	}
	offset := (page - 1) * pageSize
	return int(offset), int(pageSize)
}

//GetAggResults ..
func (s *Search) GetAggResults(aggsParams []*es.AggsParam,
	result *elastic.SearchResult) []*model.AggResult {
	for _, v := range aggsParams {
		aggResult, found := result.Aggregations.Terms("group_by_" + v.Field)
		if found {
			return s.getAggResults(aggResult)
		}
	}
	return []*model.AggResult{}
}

func (s *Search) getAggResults(aggResult *elastic.AggregationBucketKeyItems) []*model.AggResult {
	aggResults := make([]*model.AggResult, 0)
	for _, v := range aggResult.Buckets {
		aggResults = append(aggResults, &model.AggResult{
			Key:   v.Key.(string),
			Value: v.DocCount,
		})
	}
	return aggResults
}
