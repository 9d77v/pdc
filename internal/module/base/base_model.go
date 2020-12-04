package base

import (
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/utils"
	"gorm.io/gorm"
)

//Model ..
type Model struct {
	gorm.Model
	db *gorm.DB
}

//NewModel 。。
func NewModel() *Model {
	m := &Model{}
	m.db = db.GetDB()
	return m
}

//GetDB ..
func (s *Model) GetDB() *gorm.DB {
	return s.db
}

//Select 选择字段
func (s *Model) Select(fields []string, omitFields ...string) *Model {
	s.db = s.db.Select(utils.ToDBFields(fields, append([]string{"__typename"}, omitFields...)...))
	return s
}

//LeftJoin 左连接
func (s *Model) LeftJoin(query string, args ...interface{}) *Model {
	s.db = s.db.Joins("left join "+query, args...)
	return s
}

//FuzzyQuery 增加模糊查询
func (s *Model) FuzzyQuery(keyword *string, field string) *Model {
	if keyword != nil && ptrs.String(keyword) != "" {
		s.db = s.db.Where(field+" like ?", "%"+ptrs.String(keyword)+"%")
	}
	return s
}

//IDQuery id查询
func (s *Model) IDQuery(id uint, idFieldName ...string) *Model {
	s.db = s.db.Where(s.getFieldName(idFieldName...)+" = ?", id)
	return s
}

//IDArrayQuery ids查询
func (s *Model) IDArrayQuery(ids []uint, idFieldName ...string) *Model {
	if len(ids) > 0 {
		s.db = s.db.Where(s.getFieldName(idFieldName...)+" in (?)", ids)
	}
	return s
}

func (s *Model) getFieldName(idFieldName ...string) string {
	fieldName := "id"
	if len(idFieldName) > 0 {
		fieldName = idFieldName[0]
	}
	return fieldName
}

//ToUintIDs change id type from int64 to uint
func (s *Model) ToUintIDs(ids []int64) []uint {
	result := make([]uint, 0, len(ids))
	for _, id := range ids {
		result = append(result, uint(id))
	}
	return result
}

//Pagination 分页
func (s *Model) Pagination(offset, limit int) *Model {
	if limit > 0 {
		s.db = s.db.Offset(offset).Limit(limit)
	}
	return s
}

//Sort 排序
func (s *Model) Sort(sorts []*model.Sort) *Model {
	for _, v := range sorts {
		sort := " DESC"
		if v.IsAsc {
			sort = " ASC"
		}
		s.db = s.db.Order(v.Field + sort)
	}
	return s
}

//Order 排序
func (s *Model) Order(value interface{}) *Model {
	s.db = s.db.Order(value)
	return s
}

//Preload 预加载
func (s *Model) Preload(query string, args ...interface{}) *Model {
	s.db = s.db.Preload(query, args...)
	return s
}

//First 查找单挑数据并重置查询
func (s *Model) First(dest interface{}) error {
	err := s.db.First(dest).Error
	s.db = db.GetDB()
	return err
}

//Take 查找单挑数据并重置查询
func (s *Model) Take(dest interface{}) error {
	err := s.db.Take(dest).Error
	s.db = db.GetDB()
	return err
}

//Find 查找数据并重置查询
func (s *Model) Find(dest interface{}) error {
	err := s.db.Find(dest).Error
	s.db = db.GetDB()
	return err
}

//Count 查询数据总量
func (s *Model) Count(model interface{}) (total int64, err error) {
	err = s.db.Model(model).Count(&total).Error
	return
}
