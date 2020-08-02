package dtos

import (
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
)

//ToHistoryDto ...
func ToHistoryDto(m *models.History) *model.History {
	if m == nil {
		return nil
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
