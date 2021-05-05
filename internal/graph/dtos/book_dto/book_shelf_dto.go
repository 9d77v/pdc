package book_dto

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookShelfConnection(data *pb.ListBookShelfResponse) *model.BookShelfConnection {
	return &model.BookShelfConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookShelfs(data.Edges),
	}
}

func toBookShelfs(data []*pb.BookShelf) []*model.BookShelf {
	result := make([]*model.BookShelf, 0, len(data))
	for _, v := range data {
		r := toBookShelf(v)
		result = append(result, r)
	}
	return result
}

func toBookShelf(bookShelf *pb.BookShelf) *model.BookShelf {
	return &model.BookShelf{
		ID:           bookShelf.Id,
		Name:         bookShelf.Name,
		LayerNum:     int64(bookShelf.LayerNum),
		PartitionNum: int64(bookShelf.PartitionNum),
		CreatedAt:    bookShelf.CreatedAt.GetSeconds(),
		UpdatedAt:    bookShelf.UpdatedAt.GetSeconds(),
	}
}
