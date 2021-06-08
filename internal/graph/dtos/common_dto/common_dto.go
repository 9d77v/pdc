package common_dto

import (
	"context"

	"github.com/9d77v/go-pkg/ptrs"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/utils"
)

func GetSearchParam(ctx context.Context, input model.SearchParam) *base.SearchParam {
	sorts := make([]*base.Sort, 0, len(input.Sorts))
	for _, v := range input.Sorts {
		sorts = append(sorts, &base.Sort{
			Field: v.Field,
			IsAsc: v.IsAsc,
		})
	}
	return &base.SearchParam{
		QueryFields: utils.GetPreloads(ctx),
		Keyword:     ptrs.String(input.Keyword),
		Page:        ptrs.Int64(input.Page),
		PageSize:    ptrs.Int64(input.PageSize),
		Ids:         input.Ids,
		Tags:        input.Tags,
		Sorts:       sorts,
		IsRandom:    ptrs.Bool(input.IsRandom),
		IsMobile:    ptrs.Bool(input.IsMobile),
	}
}
