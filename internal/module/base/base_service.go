package base

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
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

//RecordNotExist ..
func (s Service) RecordNotExist(r Repository, id interface{}, uid ...uint) bool {
	r.IDQuery(id).Pagination(0, 1)
	if len(uid) > 0 && uid[0] != 0 {
		r.IDQuery(uid[0], "uid")
	}
	total, err := r.Count(r)
	return err != nil || total == 0
}

//GetValidIDs ..
func (s Service) GetValidIDs(r Repository, tableName string, ids []uint) []uint {
	type Temp struct {
		ID uint
	}
	temps := make([]Temp, 0, len(ids))
	fieldName := tableName + ".id"
	r.Select([]string{fieldName}).Table(tableName).IDArrayQuery(ids, fieldName).Find(&temps)
	validIDs := make([]uint, 0, len(temps))
	for _, v := range temps {
		validIDs = append(validIDs, v.ID)
	}
	return validIDs
}

//GetNewConnection get list data and total count
func (s Service) GetNewConnection(r Repository, searchParam *SearchParam,
	data interface{}, replaceFunc func(field GraphQLField) error, omitFields ...string) (total int64, err error) {
	graphqlField := ForGraphQLField(searchParam.QueryFields, "")
	if graphqlField.FieldMap["totalCount"] {
		total, err = r.Count(r)
		if err != nil {
			return
		}
	}
	offset, limit := s.GetPageInfo(searchParam)
	if graphqlField.FieldMap["edges"] {
		edgeGraphqlField := ForGraphQLField(searchParam.QueryFields, "edges.")
		r.Select(edgeGraphqlField.Fields, omitFields...)
		if searchParam.TableName != "" {
			r.IDArrayQuery(s.ToUintIDs(searchParam.Ids), searchParam.TableName+".id")
		} else {
			r.IDArrayQuery(s.ToUintIDs(searchParam.Ids))
		}
		if !searchParam.IsInfinity {
			r.Pagination(offset, limit)
		}
		r.Sort(searchParam.Sorts)
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

//GetConnection get list data and total count
func (s Service) GetConnection(ctx context.Context, r Repository, searchParam *SearchParam,
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
			IDArrayQuery(s.ToUintIDs(searchParam.Ids)).
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
func (s Service) GetPageInfo(searchParam *SearchParam) (int, int) {
	page := searchParam.Page
	pageSize := searchParam.PageSize
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

//ToUintIDs change id type from int64 to uint
func (s Service) ToUintIDs(ids []int64) []uint {
	result := make([]uint, 0, len(ids))
	for _, id := range ids {
		result = append(result, uint(id))
	}
	return result
}

//ToInt64 change id type from uint to int64
func (s Service) ToInt64(ids []uint) []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
