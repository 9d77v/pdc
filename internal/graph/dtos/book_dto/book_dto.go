package book_dto

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookConnection(data *pb.ListBookResponse) *model.BookConnection {
	return &model.BookConnection{
		TotalCount: data.TotalCount,
		Edges:      toBooks(data.Edges),
	}
}

func toBooks(data []*pb.Book) []*model.Book {
	result := make([]*model.Book, 0, len(data))
	for _, v := range data {
		r := toBook(v)
		result = append(result, r)
	}
	return result
}

func toBook(book *pb.Book) *model.Book {
	return &model.Book{
		ID:              book.Id,
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
		WordCount:       book.WordCount,
		Pricing:         book.Pricing,
		PurchasePrice:   book.PurchasePrice,
		PurchaseTime:    book.PurchaseTime.GetSeconds(),
		PurchaseSource:  book.PurchaseSource,
		BookBorrowUID:   book.BookBorrowUid,
		CreatedAt:       book.CreatedAt.GetSeconds(),
		UpdatedAt:       book.UpdatedAt.GetSeconds(),
	}
}
