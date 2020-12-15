package base

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/graph/model"
)

//Service base service
type Service struct {
}

//GetInputFields from context
func (s Service) GetInputFields(ctx context.Context) []string {
	fields := make([]string, 0)
	varibales := graphql.GetRequestContext(ctx).Variables
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	return fields
}

//GetByID Get model by id
func (s Service) GetByID(r Repository, id uint, columns []string) error {
	return r.Select(columns).IDQuery(id).First(r)
}

//GetConnection get list data and total count
func (s Service) GetConnection(ctx context.Context, r Repository, searchParam model.SearchParam,
	data interface{}, replaceFunc func(field GraphQLField) error, omitFields ...string) (total int64, err error) {
	graphqlField := NewGraphQLField(ctx, "")
	if graphqlField.FieldMap["totalCount"] {
		total, err = r.Count(r)
		if err != nil {
			return
		}
	}
	offset, limit := s.GetPageInfo(searchParam)
	if graphqlField.FieldMap["edges"] {
		edgeGraphqlField := NewGraphQLField(ctx, "edges.")
		r.Select(edgeGraphqlField.Fields, omitFields...).
			IDArrayQuery(r.ToUintIDs(searchParam.Ids)).
			Pagination(offset, limit).
			Sort(searchParam.Sorts)
		if replaceFunc != nil {
			err = replaceFunc(edgeGraphqlField)
			if err != nil {
				return
			}
		}
		err = r.Find(data)
	}
	return
}

//GetPageInfo ...
func (s Service) GetPageInfo(searchParam model.SearchParam) (int, int) {
	page := ptrs.Int64(searchParam.Page)
	pageSize := ptrs.Int64(searchParam.PageSize)
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize
	return int(offset), int(pageSize)
}