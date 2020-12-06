package base

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/utils"
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

//GetConnection get list data and total count
func (s Service) GetConnection(ctx context.Context, r Repository, searchParam model.SearchParam,
	data interface{}, replaceFunc func(edgeFieldMap map[string]bool, edgeFields []string) error,
	omitFields ...string) (total int64, err error) {
	fieldMap, _ := utils.GetFieldData(ctx, "")
	if fieldMap["totalCount"] {
		total, err = r.Count(r)
		if err != nil {
			return
		}
	}
	offset, limit := s.GetPageInfo(searchParam.Page, searchParam.PageSize)
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		r.Select(edgeFields, omitFields...).
			IDArrayQuery(r.ToUintIDs(searchParam.Ids)).
			Pagination(offset, limit).
			Sort(searchParam.Sorts)
		if replaceFunc != nil {
			err = replaceFunc(edgeFieldMap, edgeFields)
			if err != nil {
				return
			}
		}
		err = r.Find(data)
	}
	return
}

//GetPageInfo ...
func (s Service) GetPageInfo(page, pageSize *int64) (int, int) {
	offset := ptrs.Int64(page)
	limit := ptrs.Int64(pageSize)
	if offset < 1 {
		offset = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	offset = (offset - 1) * limit
	return int(offset), int(limit)
}
