package base

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"unicode"

	"gorm.io/gorm"
)

//Repository 。。
type Repository interface {
	SetDB(db *gorm.DB)
	ReSetDB()
	GetDB() *gorm.DB
	Table(name string, args ...interface{}) Repository
	SelectWithPrefix(fields []string, prefix string, omitFields ...string) Repository
	Select(fields []string, omitFields ...string) Repository
	LeftJoin(query string, args ...interface{}) Repository
	FuzzyQuery(keyword string, field string) Repository
	IDQuery(id interface{}, idFieldName ...string) Repository
	IDArrayQuery(ids interface{}, idFieldName ...string) Repository
	Where(query interface{}, args ...interface{}) Repository
	Pagination(offset, limit int) Repository
	Sort(sorts []*Sort) Repository
	Order(value interface{}) Repository
	Preload(query string, args ...interface{}) Repository
	First(dest interface{}) error
	Take(dest interface{}) error
	Find(dest interface{}) error
	Count(model interface{}) (total int64, err error)
	Create(value interface{}) error
	Updates(value interface{}) error
	Delete(value interface{}, ids []int64) error
	Begin()
	Rollback()
	Commit() error
}

//Model ..
type Model struct {
	db *gorm.DB
	tx *gorm.DB
}

//DefaultModel ..
type DefaultModel struct {
	Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

//SetDB ..
func (m *Model) SetDB(db *gorm.DB) {
	m.db = db
	m.tx = db
}

//GetDB ..
func (m *Model) GetDB() *gorm.DB {
	return m.db
}

//ReSetDB ..
func (m *Model) ReSetDB() {
	m.tx = m.db
}

//Table ..
func (m *Model) Table(name string, args ...interface{}) Repository {
	m.tx = m.tx.Table(name, args)
	return m
}

//SelectWithPrefix 选择字段带上特定前缀
func (m *Model) SelectWithPrefix(fields []string, prefix string, omitFields ...string) Repository {
	m.tx = m.tx.Select(m.toDBFieldsWithPrefix(fields, prefix,
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
			j := i + 1
			for ; j < len(s) && unicode.IsUpper(rune(s[j])); j++ {
			}
			if j == len(s) {
				j--
			}
			if i != 0 {
				newStr += "_"
			}
			newStr += strings.ToLower(string(s[i : j+1]))
			i = j
		} else {
			newStr += string(s[i])
		}
	}
	return newStr
}

//Select 选择字段
func (m *Model) Select(fields []string, omitFields ...string) Repository {
	m.tx = m.tx.Select(m.toDBFields(fields, append([]string{"__typename"}, omitFields...)...))
	return m
}

//LeftJoin 左连接
func (m *Model) LeftJoin(query string, args ...interface{}) Repository {
	m.tx = m.tx.Joins("left join "+query, args...)
	return m
}

//FuzzyQuery 增加模糊查询
func (m *Model) FuzzyQuery(keyword string, field string) Repository {
	if keyword != "" {
		m.tx = m.tx.Where(field+" like ?", "%"+keyword+"%")
	}
	return m
}

//IDQuery id查询
func (m *Model) IDQuery(id interface{}, idFieldName ...string) Repository {
	m.tx = m.tx.Where(m.getFieldName(idFieldName...)+" = ?", id)
	return m
}

//IDArrayQuery ids查询
func (m *Model) IDArrayQuery(ids interface{}, idFieldName ...string) Repository {
	switch ids.(type) {
	case []int64:
		t := ids.([]int64)
		if len(t) == 0 {
			return m
		}
	case []uint:
		t := ids.([]uint)
		if len(t) == 0 {
			return m
		}
	case []string:
		t := ids.([]string)
		if len(t) == 0 {
			return m
		}
	default:
		return m
	}
	m.tx = m.tx.Where(m.getFieldName(idFieldName...)+" in (?)", ids)
	return m
}

//Where where查询
func (m *Model) Where(query interface{}, args ...interface{}) Repository {
	m.tx = m.tx.Where(query, args...)
	return m
}

func (m *Model) getFieldName(idFieldName ...string) string {
	fieldName := "id"
	if len(idFieldName) > 0 && idFieldName[0] != "" {
		fieldName = idFieldName[0]
	}
	return fieldName
}

//Pagination 分页
func (m *Model) Pagination(offset, limit int) Repository {
	if limit > 0 {
		m.tx = m.tx.Offset(offset).Limit(limit)
	}
	return m
}

//Sort 排序
func (m *Model) Sort(sorts []*Sort) Repository {
	for _, v := range sorts {
		sort := " DESC"
		if v.IsAsc {
			sort = " ASC"
		}
		m.tx = m.tx.Order(v.Field + sort)
	}
	return m
}

//Order 排序
func (m *Model) Order(value interface{}) Repository {
	m.tx = m.tx.Order(value)
	return m
}

//Preload 预加载
func (m *Model) Preload(query string, args ...interface{}) Repository {
	m.tx = m.tx.Preload(query, args...)
	return m
}

//First 查找按主键排序第一条数据
func (m *Model) First(dest interface{}) error {
	err := m.tx.First(dest).Error
	m.ReSetDB()
	return err
}

//Take 查找单条数据
func (m *Model) Take(dest interface{}) error {
	err := m.tx.Take(dest).Error
	m.ReSetDB()
	return err
}

//Find 查找数据
func (m *Model) Find(dest interface{}) error {
	err := m.tx.Find(dest).Error
	m.ReSetDB()
	return err
}

//Count 查询数据总量
func (m *Model) Count(model interface{}) (total int64, err error) {
	err = m.tx.Model(model).Count(&total).Error
	return
}

//Create 新建数据
func (m *Model) Create(value interface{}) (err error) {
	err = m.tx.Create(value).Error
	m.ReSetDB()
	return
}

//Updates 更新数据
func (m *Model) Updates(value interface{}) (err error) {
	err = m.tx.Updates(value).Error
	m.ReSetDB()
	return
}

//Save 保存数据
func (m *Model) Save(value interface{}) (err error) {
	err = m.tx.Save(value).Error
	m.ReSetDB()
	return
}

//Delete 删除数据
func (m *Model) Delete(value interface{}, ids []int64) (err error) {
	defer func() {
		m.ReSetDB()
	}()
	if len(ids) == 0 {
		return nil
	}
	err = m.tx.Delete(value, ids).Error
	return
}

//Begin 开启事务
func (m *Model) Begin() {
	m.db = m.db.Begin()
	m.ReSetDB()
}

//Rollback 撤回事务
func (m *Model) Rollback() {
	m.tx.Rollback()
}

//Commit 提交事务
func (m *Model) Commit() (err error) {
	return m.tx.Commit().Error
}
