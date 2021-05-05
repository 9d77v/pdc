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
	BookID      uint `gorm:"uniqueIndex:book_position_uix;comment:'书id'"`
	BookShelfID uint `gorm:"comment:'书架id'"`
	Book
	BookShelf
	Layer     int8 `gorm:"comment:'层'"`
	Partition int8 `gorm:"comment:'分区'"`
}

func NewBookPositionFromPB(in *pb.CreateBookPositionRequest) *BookPosition {
	m := &BookPosition{
		BookShelfID: uint(in.BookShelfId),
		BookID:      uint(in.BookId),
		Layer:       int8(in.Layer),
		Partition:   int8(in.Partition),
	}
	m.SetDB(db.GetDB())
	return m
}

func NewBookPosition() *BookPosition {
	m := &BookPosition{}
	m.SetDB(db.GetDB())
	return m
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
		BookShelfID: int64(bookPosition.BookShelfID),
		BookShelf: &pb.BookShelf{
			Name: bookPosition.BookShelf.Name,
		},
		BookID: int64(bookPosition.BookID),
		Book: &pb.Book{
			Name: bookPosition.Book.Name,
		},
		Layer:     uint32(bookPosition.Layer),
		Partition: uint32(bookPosition.Partition),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
