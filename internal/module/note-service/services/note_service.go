package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/note-service/models"
	"github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/golang/protobuf/ptypes"
)

//NoteService ...
type NoteService struct {
	base.Service
}

//SyncNotes ...
func (s NoteService) SyncNotes(ctx context.Context,
	in *pb.SyncNotesRequest) (response *pb.SyncNotesResponse, err error) {
	response = new(pb.SyncNotesResponse)
	m := models.NewNote()
	if m.HasInvalidNotes(in) {
		err = status.Errorf(codes.InvalidArgument, "笔记格式不正确")
		return
	}
	m.Begin()
	serverNotes, err := m.GetNotesForSync(in)
	if err != nil {
		m.Rollback()
		err = status.Errorf(codes.Internal, "数据库出错：%v", err)
		return
	}

	createNotes, updateNotes, deleteNotes, clientShouldSyncNotes := m.ClassifyNotes(serverNotes, in.UnsyncedNotes, in.SyncLocal)
	if len(createNotes) > 0 {
		err = m.Create(createNotes)
		if err != nil {
			m.Rollback()
			err = status.Errorf(codes.Internal, "新增数据出错")
			return
		}
		for _, v := range createNotes {
			if v.NoteType == int8(pb.NoteType_File) {
				err = m.Create(models.NewNoteHistory(v))
				if err != nil {
					m.Rollback()
					err = status.Errorf(codes.Internal, "新增数据历史出错")
					return
				}
			}
		}
	}
	err = m.UpdateNotes(updateNotes)
	if err != nil {
		m.Rollback()
		err = status.Errorf(codes.Internal, "修改数据出错")
		return
	}
	for _, v := range deleteNotes {
		v.SetDB(m.GetDB())
		err = v.DeleteAllChildNodes()
		if err != nil {
			m.Rollback()
			err = status.Errorf(codes.Internal, "删除数据出错")
			return
		}
	}
	notes := m.GetLatestNotes(append(append(createNotes, updateNotes...), deleteNotes...))
	clientShouldSyncNotes = append(clientShouldSyncNotes, m.Sorts(notes)...)
	response.LastUpdateTime = ptypes.TimestampNow()
	response.List = m.ToNotePBs(clientShouldSyncNotes)
	err = m.Commit()
	return
}
