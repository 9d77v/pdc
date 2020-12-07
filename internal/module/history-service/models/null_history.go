package models

import (
	"errors"

	"github.com/9d77v/pdc/internal/module/base"
)

//NullHistory ...
type NullHistory struct {
}

//GetSourceID ..
func (h NullHistory) GetSourceID(subSourceID *int64) uint {
	return 0
}

//JoinSource ..
func (h NullHistory) JoinSource(r base.Repository, tableName string, sourceID uint) error {
	return errors.New("null source can not join")
}
