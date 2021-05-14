package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"github.com/golang/protobuf/ptypes"
)

const (
	OperationBorrow = iota
	OperationReturn
)

//BookBorrowReturn 书籍位置
type BookBorrowReturn struct {
	base.Model
	ID        uint `gorm:"primarykey"`
	BookID    uint `gorm:"comment:'书id'"`
	Book      Book `gorm:"ForeignKey:BookID"`
	UID       uint `gorm:"comment:'用户id'"`
	Operation int8 `gorm:"comment:'操作类型'"` //0:借，1：还
	CreatedAt time.Time
}

func NewBookBorrowReturn() *BookBorrowReturn {
	m := &BookBorrowReturn{}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *BookBorrowReturn) TableName() string {
	return db.TablePrefix() + "book_borrow_return"
}

//ToBookBorrowReturnPBs ..
func (m *BookBorrowReturn) ToBookBorrowReturnPBs(data []*BookBorrowReturn) []*pb.BookBorrowReturn {
	result := make([]*pb.BookBorrowReturn, 0, len(data))
	for _, v := range data {
		r := m.toBookBorrowReturnPB(v)
		result = append(result, r)
	}
	return result
}

func (m *BookBorrowReturn) toBookBorrowReturnPB(bookBorrowReturn *BookBorrowReturn) *pb.BookBorrowReturn {
	createdAt, _ := ptypes.TimestampProto(bookBorrowReturn.CreatedAt)
	return &pb.BookBorrowReturn{
		BookId:    int64(bookBorrowReturn.BookID),
		Uid:       int64(bookBorrowReturn.UID),
		Operation: int32(bookBorrowReturn.Operation),
		Book: &pb.Book{
			Name: bookBorrowReturn.Book.Name,
		},
		CreatedAt: createdAt,
	}
}
