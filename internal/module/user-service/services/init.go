package services

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/user-service/models"
)

func toUserDto(m *models.User, scheme string) *model.User {
	avatar := ""
	if m.Avatar != "" {
		avatar = oss.GetOSSPrefix(scheme) + m.Avatar
	}
	return &model.User{
		ID:        int64(m.ID),
		Name:      m.Name,
		Avatar:    &avatar,
		RoleID:    int64(m.RoleID),
		Gender:    int64(m.Gender),
		Color:     &m.Color,
		BirthDate: m.BirthDate.Unix(),
		IP:        &m.IP,
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
	}
}
