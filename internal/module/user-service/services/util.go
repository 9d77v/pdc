package services

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/user-service/models"
)

func (s UserService) getUsers(data []*models.User, scheme string) []*model.User {
	result := make([]*model.User, 0, len(data))
	for _, m := range data {
		r := s.getUser(m, scheme)
		result = append(result, r)
	}
	return result
}

func (s UserService) getUser(m *models.User, scheme string) *model.User {
	avatar := ""
	if m.Avatar != "" {
		avatar = oss.GetOSSPrefixByScheme(scheme) + m.Avatar
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
