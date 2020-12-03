package services

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/history-service/models"
)

func toHistoryDto(m *models.History) *model.History {
	if m == nil {
		return &model.History{}
	}
	return &model.History{
		UID:           int64(m.UID),
		SourceType:    int64(m.SourceType),
		SourceID:      int64(m.SourceID),
		SubSourceID:   int64(m.SubSourceID),
		CurrentTime:   m.CurrentTime,
		RemainingTime: m.RemainingTime,
		UpdatedAt:     m.UpdatedAt.Unix(),
	}
}
