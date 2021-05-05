package book_dto

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookBorrowReturnConnection(data *pb.ListBookBorrowReturnResponse) *model.BookBorrowReturnConnection {
	return &model.BookBorrowReturnConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookBorrowReturns(data.Edges),
	}
}

func toBookBorrowReturns(data []*pb.BookBorrowReturn) []*model.BookBorrowReturn {
	result := make([]*model.BookBorrowReturn, 0, len(data))
	for _, v := range data {
		r := toBookBorrowReturn(v)
		result = append(result, r)
	}
	return result
}

func toBookBorrowReturn(bookBorrowReturn *pb.BookBorrowReturn) *model.BookBorrowReturn {
	return &model.BookBorrowReturn{
		BookID:    bookBorrowReturn.BookId,
		UID:       bookBorrowReturn.Uid,
		Operation: int64(bookBorrowReturn.Operation),
		CreatedAt: bookBorrowReturn.CreatedAt.GetSeconds(),
	}
}
