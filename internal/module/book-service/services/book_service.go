package services

import (
	"context"
	"errors"
	"log"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/book-service/models"
	"github.com/9d77v/pdc/internal/module/book-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//BookService ..
type BookService struct {
	base.Service
}

//CreateBook ..
func (s BookService) CreateBook(ctx context.Context, in *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	resp := new(pb.CreateBookResponse)
	m := models.NewBookFromPB(in)
	err := m.Create(m)
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateBook ..
func (s BookService) UpdateBook(ctx context.Context,
	in *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	resp := &pb.UpdateBookResponse{
		Id: in.Id,
	}
	m := models.NewBook()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	updateMap := map[string]interface{}{
		"name":             in.Name,
		"desc":             in.Desc,
		"author":           &in.Author,
		"translator":       &in.Translator,
		"publishing_house": in.PublishingHouse,
		"edition":          in.Edition,
		"printed_times":    in.PrintedTimes,
		"printed_sheets":   in.PrintedSheets,
		"format":           in.Format,
		"word_count":       in.WordCount,
		"pricing":          in.Pricing,
		"purchase_price":   in.PurchasePrice,
		"purchase_time":    in.PurchaseTime.AsTime(),
		"purchase_source":  in.PurchaseSource,
	}
	if in.Cover != "" {
		updateMap["cover"] = in.Cover
	}
	return resp, m.Updates(updateMap)
}

//CreateBookshelf ..
func (s BookService) CreateBookshelf(ctx context.Context,
	in *pb.CreateBookshelfRequest) (*pb.CreateBookshelfResponse, error) {
	resp := new(pb.CreateBookshelfResponse)
	m := models.NewBookshelfFromPB(in)
	err := m.Create(m)
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateBookshelf ..
func (s BookService) UpdateBookshelf(ctx context.Context,
	in *pb.UpdateBookshelfRequest) (*pb.UpdateBookshelfResponse, error) {
	resp := &pb.UpdateBookshelfResponse{
		Id: in.Id,
	}
	m := models.NewBookshelf()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	updateMap := map[string]interface{}{
		"name": in.Name,
	}
	if in.Cover != "" {
		updateMap["cover"] = in.Cover
	}
	return resp, m.Updates(updateMap)
}

//CreateBookPosition ..
func (s BookService) CreateBookPosition(ctx context.Context,
	in *pb.CreateBookPositionRequest) (*pb.CreateBookPositionResponse, error) {
	resp := new(pb.CreateBookPositionResponse)
	m := models.NewBookPosition()
	m.Begin()
	err := m.GetLastBookPositionID(int(in.BookshelfId), int(in.Layer), int(in.Partition))
	if err != nil {
		m.Rollback()
		return resp, err
	}
	bps := make([]*models.BookPosition, len(in.BookIds))
	for i, id := range in.BookIds {
		bps[i] = &models.BookPosition{
			BookshelfID: uint(in.BookshelfId),
			BookID:      uint(id),
			Layer:       int8(in.Layer),
			Partition:   int8(in.Partition),
		}
	}
	err = m.Create(&bps)
	if err != nil {
		m.Rollback()
		return resp, err
	}
	bps[0].PrevID = m.ID
	for i := 1; i < len(bps); i++ {
		bps[i].PrevID = bps[i-1].ID
	}
	if len(bps) > 1 || bps[0].PrevID != 0 {
		err = m.Save(&bps)
		if err != nil {
			m.Rollback()
			return resp, err
		}
	}
	resp.Id = int64(m.ID)
	return resp, m.Commit()
}

//UpdateBookPosition ..
func (s BookService) UpdateBookPosition(ctx context.Context,
	in *pb.UpdateBookPositionRequest) (*pb.UpdateBookPositionResponse, error) {
	resp := &pb.UpdateBookPositionResponse{Id: in.Id}
	m := models.NewBookPosition()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"book_shelf_id": in.BookshelfId,
		"layer":         in.Layer,
		"partition":     in.Partition,
	})
}

//RemoveBookPosition ..
func (s BookService) RemoveBookPosition(ctx context.Context,
	in *pb.RemoveBookPositionRequest) (*pb.RemoveBookPositionResponse, error) {
	resp := &pb.RemoveBookPositionResponse{Id: in.Id}
	m := models.NewBookPosition()
	m.Begin()
	err := s.GetByID(m, uint(in.Id), []string{"id", "prev_id"})
	if err != nil {
		m.Rollback()
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	prevID := m.PrevID
	m.ID = 0
	err = m.Where("prev_id=?", in.Id).First(m)
	log.Println(m)
	if err == nil {
		m.Table(m.TableName()).Where("id=?", m.ID).Updates(map[string]interface{}{
			"prev_id": prevID,
		})
	}
	m.DeleteByID(in.Id)
	return resp, m.Commit()
}

//BorrowBook ..
func (s BookService) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	resp := new(pb.BorrowBookResponse)
	book := models.NewBook()
	book.Begin()
	if err := s.GetByID(book, uint(in.BookId), []string{"id", "borrow_id"}); err != nil {
		book.Rollback()
		return nil, err
	}
	if book.BookBorrowUID != 0 {
		book.Rollback()
		return nil, errors.New("book is borrowed")
	}
	m := models.NewBookBorrowReturn()
	m.BookID = uint(in.BookId)
	m.UID = uint(in.Uid)
	m.Operation = models.OperationBorrow
	err := book.Create(m)
	if err != nil {
		book.Rollback()
		return nil, errors.New("create book operation history failed")
	}
	err = book.Updates(map[string]interface{}{
		"borrow_id": in.Uid,
	})
	if err != nil {
		book.Rollback()
		return nil, errors.New("update book borrow_id failed")
	}
	resp.Id = int64(m.ID)
	return resp, book.Commit()
}

//ReturnBook ..
func (s BookService) ReturnBook(ctx context.Context, in *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	resp := new(pb.ReturnBookResponse)
	book := models.NewBook()
	book.Begin()
	if err := s.GetByID(book, uint(in.BookId), []string{"id", "borrow_id"}); err != nil {
		book.Rollback()
		return nil, err
	}
	if book.BookBorrowUID == 0 {
		book.Rollback()
		return nil, errors.New("book is not borrowed")
	}
	m := models.NewBookBorrowReturn()
	m.BookID = uint(in.BookId)
	m.UID = uint(in.Uid)
	m.Operation = models.OperationReturn
	err := book.Create(m)
	if err != nil {
		book.Rollback()
		return nil, errors.New("create book operation history failed")
	}
	err = book.Updates(map[string]interface{}{
		"borrow_id": 0,
	})
	if err != nil {
		book.Rollback()
		return nil, errors.New("update book borrow_id failed")
	}
	resp.Id = int64(m.ID)
	return resp, err
}

//ListBook ..
func (s BookService) ListBook(ctx context.Context, in *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	resp := new(pb.ListBookResponse)
	m := models.NewBook()
	m.FuzzyQuery(in.SearchParam.Keyword, "name")
	if ptrs.Bool(in.FilterBooksInBookPositions) {
		m.Where("id NOT in (select book_id from " + new(models.BookPosition).TableName() + ")")
	}
	data := make([]*models.Book, 0)
	total, err := s.GetNewConnection(m, in.SearchParam, &data, nil)
	resp.TotalCount = total
	resp.Edges = m.ToBookPBs(data)
	return resp, err
}

//ListBookshelf ..
func (s BookService) ListBookshelf(ctx context.Context, in *pb.ListBookshelfRequest) (*pb.ListBookshelfResponse, error) {
	resp := new(pb.ListBookshelfResponse)
	m := models.NewBookshelf()
	m.FuzzyQuery(in.SearchParam.Keyword, "name")
	data := make([]*models.Bookshelf, 0)
	total, err := s.GetNewConnection(m, in.SearchParam, &data, nil)
	resp.TotalCount = total
	resp.Edges = m.ToBookshelfPBs(data)
	return resp, err
}

//ListBookPosition ..
func (s BookService) ListBookPosition(ctx context.Context, in *pb.ListBookPositionRequest) (*pb.ListBookPositionResponse, error) {
	resp := new(pb.ListBookPositionResponse)
	m := models.NewBookPosition()
	m.FuzzyQuery(in.SearchParam.Keyword, "name")
	if in.BookId != nil {
		m.IDQuery(*in.BookId, "book_id")
	}
	if in.BookshelfId != nil {
		m.IDQuery(*in.BookshelfId, "bookshelf_id")
	}
	data := make([]*models.BookPosition, 0)
	replaceFunc := func(field base.GraphQLField) error {
		if field.FieldMap["book"] {
			m.Preload("Book")
		}
		if field.FieldMap["bookshelf"] {
			m.Preload("Bookshelf")
		}
		return nil
	}
	omitFields := []string{"book", "bookshelf"}
	total, err := s.GetNewConnection(m, in.SearchParam, &data, replaceFunc, omitFields...)
	resp.TotalCount = total
	resp.Edges = m.ToBookPositionPBs(data)
	return resp, err
}

//ListBookBorrowReturn ..
func (s BookService) ListBookBorrowReturn(ctx context.Context, in *pb.ListBookBorrowReturnRequest) (*pb.ListBookBorrowReturnResponse, error) {
	resp := new(pb.ListBookBorrowReturnResponse)
	m := models.NewBookBorrowReturn()
	m.FuzzyQuery(in.SearchParam.Keyword, "name")
	if in.BookID != nil {
		m.IDQuery(*in.BookID, "book_id")
	}
	data := make([]*models.BookBorrowReturn, 0)
	replaceFunc := func(field base.GraphQLField) error {
		if field.FieldMap["book"] {
			m.Preload("Book")
		}
		if field.FieldMap["bookShelf"] {
			m.Preload("Bookshelf")
		}
		return nil
	}
	omitFields := []string{""}
	total, err := s.GetNewConnection(m, in.SearchParam, &data, replaceFunc, omitFields...)
	resp.TotalCount = total
	resp.Edges = m.ToBookBorrowReturnPBs(data)
	return resp, err
}
