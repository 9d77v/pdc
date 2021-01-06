package chmodels

import (
	"time"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/module/base"
)

//HistoryLog 历史记录
type HistoryLog struct {
	base.Model
	SourceType    uint8
	UID           uint64
	SourceID      uint64
	SubSourceID   uint64
	ServerTs      time.Time
	ClientTs      time.Time
	Platform      string
	CurrentTime   float64
	RemainingTime float64
	Duration      float64
}

//NewHistoryLog ..
func NewHistoryLog() *HistoryLog {
	m := &HistoryLog{}
	m.SetDB(clickhouse.GetDB())
	return m
}
