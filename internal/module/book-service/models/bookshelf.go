package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"github.com/golang/protobuf/ptypes"
)

//Bookshelf 书架
type Bookshelf struct {
	base.DefaultModel
	Name         string `gorm:"size:100;NOT NULL;comment:'书架名'"`
	Cover        string `gorm:"size:200;NOT NULL;comment:'图片'"`
	LayerNum     int8   `gorm:"comment:'层数'"`
	PartitionNum int8   `gorm:"comment:'分区数'"`
}

//NewBookshelf ..
func NewBookshelf() *Bookshelf {
	m := &Bookshelf{}
	m.SetDB(db.GetDB())
	return m
}

func NewBookshelfFromPB(in *pb.CreateBookshelfRequest) *Bookshelf {
	m := &Bookshelf{
		Name:         in.Name,
		Cover:        in.Cover,
		LayerNum:     int8(in.LayerNum),
		PartitionNum: int8(in.PartitionNum),
	}
	m.SetDB(db.GetDB())
	return m
}

//ToBookshelfPBs ..
func (m *Bookshelf) ToBookshelfPBs(data []*Bookshelf) []*pb.Bookshelf {
	result := make([]*pb.Bookshelf, 0, len(data))
	for _, v := range data {
		r := m.toBookshelfPB(v)
		result = append(result, r)
	}
	return result
}

func (m *Bookshelf) toBookshelfPB(bookshelf *Bookshelf) *pb.Bookshelf {
	createdAt, _ := ptypes.TimestampProto(bookshelf.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(bookshelf.UpdatedAt)
	return &pb.Bookshelf{
		Id:           int64(bookshelf.ID),
		Name:         bookshelf.Name,
		Cover:        bookshelf.Cover,
		LayerNum:     int32(bookshelf.LayerNum),
		PartitionNum: int32(bookshelf.PartitionNum),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
