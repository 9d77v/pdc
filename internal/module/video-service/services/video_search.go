package services

import (
	"context"
	"log"
	"strings"

	elastic "github.com/olivere/elastic/v7"

	es "github.com/9d77v/go-lib/clients/elastic/v7"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/elasticsearch"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
)

type videoSearch struct {
	*base.Search
}

func newVideoSearch(ctx context.Context, searchParam model.SearchParam,
	scheme string) *videoSearch {
	return &videoSearch{
		Search: base.NewSearch(ctx, searchParam, scheme),
	}
}

func (s *videoSearch) execute() (*model.VideoIndexConnection, error) {
	result := new(model.VideoIndexConnection)
	keyword := strings.ReplaceAll(ptrs.String(s.SearchParam.Keyword), " ", "")
	if keyword != "" {
		s.BoolQuery.Must(newKeywordQuery(keyword))
	}
	s.Filter()
	offset, limit := s.GetPageInfo()
	searchService := elasticsearch.GetClient().Search().
		Index(elasticsearch.AliasVideo).
		Query(s.BoolQuery).
		From(offset).
		Size(limit)
	aggsParams := []*es.AggsParam{
		{Field: "tags", Size: 50},
	}
	field := base.NewGraphQLField(s.Ctx, "")
	if field.FieldMap["aggResults"] {
		searchService = es.Aggs(searchService, aggsParams...)
	}
	if field.FieldMap["edges"] {
		s.sort(searchService)
	} else {
		searchService.FetchSource(false)
	}
	searchResult, err := searchService.Do(s.Ctx)
	if err != nil {
		log.Println("err:", err)
		return result, nil
	}
	result.TotalCount = searchResult.TotalHits()
	if field.FieldMap["edges"] {
		result.Edges = s.GetEdges(searchResult, s.Scheme)
	}
	if field.FieldMap["aggResults"] {
		result.AggResults = s.GetAggResults(aggsParams, searchResult)
	}
	return result, nil
}

func newKeywordQuery(keyword string) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(keyword, []string{
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
}

func (s *videoSearch) sort(searchService *elastic.SearchService) {
	if ptrs.Bool(s.SearchParam.IsRandom) {
		searchService.SortBy(elastic.NewScriptSort(elastic.NewScript("Math.random()"), "number").Order(true))
	} else {
		searchService.Sort("_score", false).
			Sort("title.keyword", true)
	}
}
