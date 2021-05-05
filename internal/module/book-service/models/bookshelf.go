package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"github.com/golang/protobuf/ptypes"
)

//BookShelf 书架
type BookShelf struct {
	base.DefaultModel
	Name         string `gorm:"size:100;NOT NULL;comment:'书架名'"`
	LayerNum     int8   `gorm:"comment:'层数'"`
	PartitionNum int8   `gorm:"comment:'分区数'"`
}

//NewBookShelf ..
func NewBookShelf() *BookShelf {
	m := &BookShelf{}
	m.SetDB(db.GetDB())
	return m
}

func NewBookShelfFromPB(in *pb.CreateBookShelfRequest) *BookShelf {
	m := &BookShelf{
		Name:         in.Name,
		LayerNum:     int8(in.LayerNum),
		PartitionNum: int8(in.PartitionNum),
	}
	m.SetDB(db.GetDB())
	return m
}

//ToBookShelfPBs ..
func (m *BookShelf) ToBookShelfPBs(data []*BookShelf) []*pb.BookShelf {
	result := make([]*pb.BookShelf, 0, len(data))
	for _, v := range data {
		r := m.toBookShelfPB(v)
		result = append(result, r)
	}
	return result
}

func (m *BookShelf) toBookShelfPB(bookShelf *BookShelf) *pb.BookShelf {
	createdAt, _ := ptypes.TimestampProto(bookShelf.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(bookShelf.UpdatedAt)
	return &pb.BookShelf{
		Id:           int64(bookShelf.ID),
		Name:         bookShelf.Name,
		LayerNum:     int32(bookShelf.LayerNum),
		PartitionNum: int32(bookShelf.PartitionNum),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
