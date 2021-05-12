package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"github.com/golang/protobuf/ptypes"
)

//BookPosition 书籍位置
type BookPosition struct {
	base.DefaultModel
	BookID      uint      `gorm:"uniqueIndex:book_position_uix;comment:'书id'"`
	BookshelfID uint      `gorm:"comment:'书架id'"`
	Book        Book      `gorm:"ForeignKey:BookID"`
	Bookshelf   Bookshelf `gorm:"ForeignKey:BookshelfID"`
	Layer       int8      `gorm:"comment:'层'"`
	Partition   int8      `gorm:"comment:'分区'"`
	PrevID      uint      `gorm:"comment:'左侧位置id'"`
}

func NewBookPosition() *BookPosition {
	m := &BookPosition{}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *BookPosition) TableName() string {
	return db.TablePrefix() + "book_position"
}

//ToBookPositionPBs ..
func (m *BookPosition) ToBookPositionPBs(data []*BookPosition) []*pb.BookPosition {
	result := make([]*pb.BookPosition, 0, len(data))
	for _, v := range data {
		r := m.toBookPositionPB(v)
		result = append(result, r)
	}
	return result
}

func (m *BookPosition) toBookPositionPB(bookPosition *BookPosition) *pb.BookPosition {
	createdAt, _ := ptypes.TimestampProto(bookPosition.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(bookPosition.UpdatedAt)
	return &pb.BookPosition{
		Id:          int64(bookPosition.ID),
		BookshelfId: int64(bookPosition.BookshelfID),
		Bookshelf: &pb.Bookshelf{
			Name:  bookPosition.Bookshelf.Name,
			Cover: bookPosition.Bookshelf.Cover,
		},
		BookID: int64(bookPosition.BookID),
		Book: &pb.Book{
			Name:  bookPosition.Book.Name,
			Cover: bookPosition.Book.Cover,
		},
		Layer:     uint32(bookPosition.Layer),
		Partition: uint32(bookPosition.Partition),
		PrevId:    int64(bookPosition.PrevID),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

//GetLastBookPosition 获取某个书架某行某分区的最后一本书id
func (m *BookPosition) GetLastBookPositionID(bookshelfID, layer, partiton int) error {
	return db.GetDB().Raw(`WITH RECURSIVE "bp" (id,prev_id) as
	(
	 select id,prev_id from "pdc_book_position" where bookshelf_id=? and layer=? and partition=? and prev_id=0
	 union all
	 select e2.id,e2.prev_id from "pdc_book_position" e2,"bp" e3 where e3.id=e2.prev_id
	)
	select id,ROW_NUMBER() OVER() as rowId from bp order by rowId desc limit 1`, bookshelfID, layer, partiton).Scan(m).Error
}

//DeleteByID 获取某个书架某行某分区的最后一本书id
func (m *BookPosition) DeleteByID(id int64) error {
	m.ID = uint(id)
	return db.GetDB().Unscoped().Delete(m).Error
}
