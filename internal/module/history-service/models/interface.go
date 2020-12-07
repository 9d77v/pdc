package models

import (
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
)

//Histroyer ...
type Histroyer interface {
	GetSourceID(subSourceID *int64) uint
	JoinSource(r base.Repository, tableName string, sourceID uint) error
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
