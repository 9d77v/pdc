package book_dto

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
)

func GetBookshelfConnection(data *pb.ListBookshelfResponse, scheme string) *model.BookshelfConnection {
	return &model.BookshelfConnection{
		TotalCount: data.TotalCount,
		Edges:      toBookshelfs(data.Edges, scheme),
	}
}

func toBookshelfs(data []*pb.Bookshelf, scheme string) []*model.Bookshelf {
	result := make([]*model.Bookshelf, 0, len(data))
	for _, v := range data {
		r := toBookshelf(v, scheme)
		result = append(result, r)
	}
	return result
}

func toBookshelf(bookShelf *pb.Bookshelf, scheme string) *model.Bookshelf {
	cover := ""
	if bookShelf.Cover != "" {
		cover = oss.GetOSSPrefixByScheme(scheme) + bookShelf.Cover
	}
	return &model.Bookshelf{
		ID:           bookShelf.Id,
		Name:         bookShelf.Name,
		Cover:        cover,
		LayerNum:     int64(bookShelf.LayerNum),
		PartitionNum: int64(bookShelf.PartitionNum),
		CreatedAt:    bookShelf.CreatedAt.GetSeconds(),
		UpdatedAt:    bookShelf.UpdatedAt.GetSeconds(),
	}
}
