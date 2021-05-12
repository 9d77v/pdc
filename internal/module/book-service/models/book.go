package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
)

//Book 书
type Book struct {
	base.DefaultModel
	Isbn            string         `gorm:"size:32;NOT NULL;comment:'isbn'"`
	Name            string         `gorm:"size:100;NOT NULL;comment:'书名'"`
	Desc            string         `gorm:"size:1000;NOT NULL;comment:'简介'"`
	Cover           string         `gorm:"size:200;NOT NULL;comment:'封面'"`
	Author          pq.StringArray `gorm:"type:varchar(100)[];comment:'作者'"`
	Translator      pq.StringArray `gorm:"type:varchar(100)[];comment:'译者'"`
	PublishingHouse string         `gorm:"size:100;NOT NULL;comment:'出版社'"`
	Edition         string         `gorm:"size:32;NOT NULL;comment:'版次'"`
	PrintedTimes    string         `gorm:"size:32;NOT NULL;comment:'印次'"`
	PrintedSheets   string         `gorm:"size:32;NOT NULL;comment:'印张'"`
	Format          string         `gorm:"size:32;NOT NULL;comment:'开本'"`
	WordCount       int            `gorm:"comment:'字数'"`
	Pricing         float32        `gorm:"money;NOT NULL;comment:'定价'"`
	PurchasePrice   float32        `gorm:"money;NOT NULL;comment:'购买价'"`
	PurchaseTime    time.Time      `gorm:"comment:'购买时间'"`
	PurchaseSource  string         `gorm:"size:100;NOT NULL;comment:'购买途径'"`
	BookBorrowUID   uint           `gorm:"comment:'借书人uid'"`
}

func NewBook() *Book {
	m := &Book{}
	m.SetDB(db.GetDB())
	return m
}

func NewBookFromPB(in *pb.CreateBookRequest) *Book {
	m := &Book{
		Isbn:            in.Isbn,
		Name:            in.Name,
		Desc:            in.Desc,
		Cover:           in.Cover,
		Author:          in.Author,
		Translator:      in.Translator,
		PublishingHouse: in.PublishingHouse,
		Edition:         in.Edition,
		PrintedTimes:    in.PrintedTimes,
		PrintedSheets:   in.PrintedSheets,
		Format:          in.Format,
		WordCount:       int(in.WordCount),
		Pricing:         float32(in.Pricing),
		PurchasePrice:   float32(in.PurchasePrice),
		PurchaseTime:    in.PurchaseTime.AsTime(),
		PurchaseSource:  in.PurchaseSource,
	}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *Book) TableName() string {
	return db.TablePrefix() + "book"
}

//ToBookPBs ..
func (m *Book) ToBookPBs(data []*Book) []*pb.Book {
	result := make([]*pb.Book, 0, len(data))
	for _, v := range data {
		r := m.toBookPB(v)
		result = append(result, r)
	}
	return result
}

func (m *Book) toBookPB(book *Book) *pb.Book {
	createdAt, _ := ptypes.TimestampProto(book.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(book.UpdatedAt)
	purchaseTime, _ := ptypes.TimestampProto(book.PurchaseTime)
	return &pb.Book{
		Id:              int64(book.ID),
		Isbn:            book.Isbn,
		Name:            book.Name,
		Desc:            book.Desc,
		Cover:           book.Cover,
		Author:          book.Author,
		Translator:      book.Translator,
		PublishingHouse: book.PublishingHouse,
		Edition:         book.Edition,
		PrintedTimes:    book.PrintedTimes,
		PrintedSheets:   book.PrintedSheets,
		Format:          book.Format,
		WordCount:       int64(book.WordCount),
		Pricing:         float64(book.Pricing),
		PurchasePrice:   float64(book.PurchasePrice),
		PurchaseTime:    purchaseTime,
		PurchaseSource:  book.PurchaseSource,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
