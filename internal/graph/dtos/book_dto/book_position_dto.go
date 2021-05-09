package book_dto

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookPositionConnection(data *pb.ListBookPositionResponse) *model.BookPositionConnection {
	return &model.BookPositionConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookPositions(data.Edges),
	}
}

func toBookPositions(data []*pb.BookPosition) []*model.BookPosition {
	result := make([]*model.BookPosition, 0, len(data))
	for _, v := range data {
		r := toBookPosition(v)
		result = append(result, r)
	}
	return result
}

func toBookPosition(bookPosition *pb.BookPosition) *model.BookPosition {
	return &model.BookPosition{
		BookshelfID: bookPosition.BookshelfId,
		BookID:      bookPosition.BookID,
		Layer:       int64(bookPosition.Layer),
		Partition:   int64(bookPosition.Partition),
		CreatedAt:   bookPosition.CreatedAt.GetSeconds(),
		UpdatedAt:   bookPosition.UpdatedAt.GetSeconds(),
	}
}
