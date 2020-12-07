package base

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/model"
	"gorm.io/gorm"
)

//Repository 。。
type Repository interface {
	GetDB() *gorm.DB
	Table(name string, args ...interface{}) Repository
	SelectWithPrefix(fields []string, prefix string, omitFields ...string) Repository
	Select(fields []string, omitFields ...string) Repository
	LeftJoin(query string, args ...interface{}) Repository
	FuzzyQuery(keyword *string, field string) Repository
	IDQuery(id uint, idFieldName ...string) Repository
	IDArrayQuery(ids []uint, idFieldName ...string) Repository
	Where(query interface{}, args ...interface{}) Repository
	ToUintIDs(ids []int64) []uint
	Pagination(offset, limit int) Repository
	Sort(sorts []*model.Sort) Repository
	Order(value interface{}) Repository
	Preload(query string, args ...interface{}) Repository
	First(dest interface{}) error
	Take(dest interface{}) error
	Find(dest interface{}) error
	Count(model interface{}) (total int64, err error)
}

//Model ..
type Model struct {
	db *gorm.DB
}

//NewModel 。。
func NewModel() *Model {
	m := &Model{}
	m.db = db.GetDB()
	return m
}

//DefaultModel ..
type DefaultModel struct {
	*Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

//NewDefaultModel 。。
func NewDefaultModel() DefaultModel {
	m := DefaultModel{}
	m.Model = NewModel()
	return m
}

//GetDB ..
func (m *Model) GetDB() *gorm.DB {
	return m.db
}

//Table ..
func (m *Model) Table(name string, args ...interface{}) Repository {
	m.db = m.db.Table(name, args)
	return m
}

//SelectWithPrefix 选择字段带上特定前缀
func (m *Model) SelectWithPrefix(fields []string, prefix string, omitFields ...string) Repository {
	m.db = m.db.Select(m.toDBFieldsWithPrefix(fields, prefix,
		append([]string{"__typename"}, omitFields...)...))
	return m
}

func (m *Model) toDBFieldsWithPrefix(fields []string, prefix string, omitFields ...string) []string {
	newFields := make([]string, 0, len(fields))
	for _, v := range m.toDBFields(fields, append([]string{"__typename"}, omitFields...)...) {
		newFields = append(newFields, prefix+v)
	}
	return newFields
}

//ToDBFields calculate slect fields by input fields
func (m *Model) toDBFields(fields []string, omitFields ...string) []string {
	dbFields := make([]string, 0)
	omitFieldMap := make(map[string]bool)
	for _, v := range omitFields {
		omitFieldMap[v] = true
	}
	for _, v := range fields {
		if !omitFieldMap[v] {
			if strings.Contains(strings.ToLower(v), "price") {
				v = fmt.Sprintf("\"%s\"::money::numeric::float8", m.camelToSnack(v))
			} else if strings.Contains(v, ".") || strings.Contains(v, " ") {
			} else {
				v = fmt.Sprintf("\"%s\"", m.camelToSnack(v))
			}
			dbFields = append(dbFields, v)
		}
	}
	return dbFields
}

func (m *Model) camelToSnack(s string) string {
	newStr := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			newStr += "_" + strings.ToLower(string(s[i]))
		} else {
			newStr += string(s[i])
		}
	}
	newStr = strings.ReplaceAll(newStr, "_i_d", "_id")
	return strings.ReplaceAll(newStr, "_u_r_l", "_url")
}

//Select 选择字段
func (m *Model) Select(fields []string, omitFields ...string) Repository {
	m.db = m.db.Select(m.toDBFields(fields, append([]string{"__typename"}, omitFields...)...))
	return m
}

//LeftJoin 左连接
func (m *Model) LeftJoin(query string, args ...interface{}) Repository {
	m.db = m.db.Joins("left join "+query, args...)
	return m
}

//FuzzyQuery 增加模糊查询
func (m *Model) FuzzyQuery(keyword *string, field string) Repository {
	if keyword != nil && ptrs.String(keyword) != "" {
		m.db = m.db.Where(field+" like ?", "%"+ptrs.String(keyword)+"%")
	}
	return m
}

//IDQuery id查询
func (m *Model) IDQuery(id uint, idFieldName ...string) Repository {
	m.db = m.db.Where(m.getFieldName(idFieldName...)+" = ?", id)
	return m
}

//IDArrayQuery ids查询
func (m *Model) IDArrayQuery(ids []uint, idFieldName ...string) Repository {
	if len(ids) > 0 {
		m.db = m.db.Where(m.getFieldName(idFieldName...)+" in (?)", ids)
	}
	return m
}

//Where where查询
func (m *Model) Where(query interface{}, args ...interface{}) Repository {
	m.db = m.db.Where(query, args...)
	return m
}

func (m *Model) getFieldName(idFieldName ...string) string {
	fieldName := "id"
	if len(idFieldName) > 0 {
		fieldName = idFieldName[0]
	}
	return fieldName
}

//ToUintIDs change id type from int64 to uint
func (m *Model) ToUintIDs(ids []int64) []uint {
	result := make([]uint, 0, len(ids))
	for _, id := range ids {
		result = append(result, uint(id))
	}
	return result
}

//Pagination 分页
func (m *Model) Pagination(offset, limit int) Repository {
	if limit > 0 {
		m.db = m.db.Offset(offset).Limit(limit)
	}
	return m
}

//Sort 排序
func (m *Model) Sort(sorts []*model.Sort) Repository {
	for _, v := range sorts {
		sort := " DESC"
		if v.IsAsc {
			sort = " ASC"
		}
		m.db = m.db.Order(v.Field + sort)
	}
	return m
}

//Order 排序
func (m *Model) Order(value interface{}) Repository {
	m.db = m.db.Order(value)
	return m
}

//Preload 预加载
func (m *Model) Preload(query string, args ...interface{}) Repository {
	m.db = m.db.Preload(query, args...)
	return m
}

//First 查找单挑数据并重置查询
func (m *Model) First(dest interface{}) error {
	err := m.db.First(dest).Error
	m.db = db.GetDB()
	return err
}

//Take 查找单挑数据并重置查询
func (m *Model) Take(dest interface{}) error {
	err := m.db.Take(dest).Error
	m.db = db.GetDB()
	return err
}

//Find 查找数据并重置查询
func (m *Model) Find(dest interface{}) error {
	err := m.db.Find(dest).Error
	m.db = db.GetDB()
	return err
}

//Count 查询数据总量
func (m *Model) Count(model interface{}) (total int64, err error) {
	err = m.db.Model(model).Count(&total).Error
	return
}
