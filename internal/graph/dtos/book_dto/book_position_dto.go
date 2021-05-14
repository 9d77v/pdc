package book_dto

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookPositionConnection(data *pb.ListBookPositionResponse, scheme string) *model.BookPositionConnection {
	return &model.BookPositionConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookPositions(data.Edges, scheme),
	}
}

func toBookPositions(data []*pb.BookPosition, scheme string) []*model.BookPosition {
	result := make([]*model.BookPosition, 0, len(data))
	for _, v := range data {
		r := toBookPosition(v, scheme)
		result = append(result, r)
	}
	return result
}

func toBookPosition(bookPosition *pb.BookPosition, scheme string) *model.BookPosition {
	book := new(model.Book)
	if bookPosition.Book != nil {
		book.Name = bookPosition.Book.Name
		cover := ""
		if bookPosition.Book.Cover != "" {
			cover = oss.GetOSSPrefixByScheme(scheme) + bookPosition.Book.Cover
		}
		book.Cover = cover
	}
	bookshelf := new(model.Bookshelf)
	if bookPosition.Bookshelf != nil {
		bookshelf.Name = bookPosition.Bookshelf.Name
		cover := ""
		if bookPosition.Bookshelf.Cover != "" {
			cover = oss.GetOSSPrefixByScheme(scheme) + bookPosition.Bookshelf.Cover
		}
		bookshelf.Cover = cover
	}
	return &model.BookPosition{
		ID:          bookPosition.Id,
		BookshelfID: bookPosition.BookshelfId,
		BookID:      bookPosition.BookID,
		Book:        book,
		Bookshelf:   bookshelf,
		Layer:       int64(bookPosition.Layer),
		PrevID:      int64(bookPosition.PrevId),
		Partition:   int64(bookPosition.Partition),
		CreatedAt:   bookPosition.CreatedAt.GetSeconds(),
		UpdatedAt:   bookPosition.UpdatedAt.GetSeconds(),
	}
}
