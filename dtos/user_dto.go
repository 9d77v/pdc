package dtos

import (
	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
)

//ToUserDto ...
func ToUserDto(m *models.User) *model.User {
	avatar := ""
	if m.Avatar != "" {
		avatar = models.OssPrefix + m.Avatar
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
