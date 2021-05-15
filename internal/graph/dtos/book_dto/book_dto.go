package book_dto

import (
	"strconv"

	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookConnection(data *pb.ListBookResponse, scheme string) *model.BookConnection {
	return &model.BookConnection{
		TotalCount: data.TotalCount,
		Edges:      toBooks(data.Edges, scheme),
	}
}

func toBooks(data []*pb.Book, scheme string) []*model.Book {
	result := make([]*model.Book, 0, len(data))
	for _, v := range data {
		r := toBook(v, scheme)
		result = append(result, r)
	}
	return result
}

func toBook(book *pb.Book, scheme string) *model.Book {
	cover := ""
	if book.Cover != "" {
		cover = oss.GetOSSPrefixByScheme(scheme) + book.Cover
	}
	return &model.Book{
		ID:              book.Id,
		Isbn:            book.Isbn,
		Name:            book.Name,
		Desc:            book.Desc,
		Cover:           cover,
		Author:          book.Author,
		Translator:      book.Translator,
		PublishingHouse: book.PublishingHouse,
		Edition:         book.Edition,
		PrintedTimes:    book.PrintedTimes,
		PrintedSheets:   book.PrintedSheets,
		Format:          book.Format,
		WordCount:       book.WordCount,
		Pricing:         strconv.FormatFloat(book.Pricing, 'f', 2, 64),
		PurchasePrice:   strconv.FormatFloat(book.PurchasePrice, 'f', 2, 64),
		PurchaseTime:    book.PurchaseTime.GetSeconds(),
		PurchaseSource:  book.PurchaseSource,
		BookBorrowUID:   book.BookBorrowUid,
		CreatedAt:       book.CreatedAt.GetSeconds(),
		UpdatedAt:       book.UpdatedAt.GetSeconds(),
	}
}

func GetBookIndexConnection(data *pb.SearchBookResponse, scheme string) *model.BookIndexConnection {
	return &model.BookIndexConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookIndexs(data.Edges, scheme),
	}
}

func toBookIndexs(data []*pb.BookIndex, scheme string) []*model.BookIndex {
	result := make([]*model.BookIndex, 0, len(data))
	for _, v := range data {
		r := toBookIndex(v, scheme)
		result = append(result, r)
	}
	return result
}

func toBookIndex(book *pb.BookIndex, scheme string) *model.BookIndex {
	cover := ""
	if book.Cover != "" {
		cover = oss.GetOSSPrefixByScheme(scheme) + book.Cover
	}
	return &model.BookIndex{
		ID:              book.Id,
		Isbn:            book.Isbn,
		Name:            book.Name,
		Desc:            book.Desc,
		Cover:           cover,
		Author:          book.Author,
		Translator:      book.Translator,
		PublishingHouse: book.PublishingHouse,
		Edition:         book.Edition,
		PrintedTimes:    book.PrintedTimes,
		PrintedSheets:   book.PrintedSheets,
		Format:          book.Format,
		WordCount:       book.WordCount,
		Pricing:         strconv.FormatFloat(book.Pricing, 'f', 2, 64),
	}
}
