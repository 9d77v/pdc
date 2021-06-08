package models

import (
	"context"

	"github.com/9d77v/go-pkg/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
)

//Histroyer ...
type Histroyer interface {
	GetSourceID(subSourceID *int64) uint
	JoinSource(r base.Repository, tableName string, sourceID uint) error
	GetStatistic(ctx context.Context, uid uint) [][]float64
}

//CreateHistory retrun history depends on sourcetype
func CreateHistory(sourceType *int64) Histroyer {
	switch ptrs.Int64(sourceType) {
	case HistoryVideo:
		return VideoHistory{}
	default:
		return NullHistory{}
	}
}
