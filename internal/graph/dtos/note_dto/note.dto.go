package note_dto

import (
	"time"

	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/golang/protobuf/ptypes"
)

func GetNotes(data []*model.NewNote) []*pb.Note {
	result := make([]*pb.Note, 0, len(data))
	for _, v := range data {
		r := getNote(v)
		result = append(result, r)
	}
	return result
}
func getNote(m *model.NewNote) *pb.Note {
	createdAt, _ := ptypes.TimestampProto(time.Unix(m.CreatedAt, 0))
	updatedAt, _ := ptypes.TimestampProto(time.Unix(m.UpdatedAt, 0))
	return &pb.Note{
		Id:        m.ID,
		ParentId:  m.ParentID,
		Uid:       int64(m.UID),
		NoteType:  pb.NoteType(m.NoteType),
		Level:     int32(m.Level),
		Title:     m.Title,
		Color:     m.Color,
		State:     pb.NoteState(m.State),
		Version:   int32(m.Version),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Content:   m.Content,
		Tags:      m.Tags,
		Sha1:      m.Sha1,
	}
}

func GetSyncNotesResponseConnection(data *pb.SyncNotesResponse) *model.SyncNotesResponse {
	return &model.SyncNotesResponse{
		LastUpdateTime: data.LastUpdateTime.Seconds,
		List:           toNotes(data.List),
	}
}

func toNotes(data []*pb.Note) []*model.Note {
	result := make([]*model.Note, 0, len(data))
	for _, v := range data {
		r := toNote(v)
		result = append(result, r)
	}
	return result
}

func toNote(m *pb.Note) *model.Note {
	return &model.Note{
		ID:        m.Id,
		ParentID:  m.ParentId,
		UID:       m.Uid,
		NoteType:  int64(m.NoteType),
		Level:     int64(m.Level),
		Title:     m.Title,
		Color:     m.Color,
		State:     int64(m.State),
		Version:   int64(m.Version),
		CreatedAt: m.CreatedAt.Seconds,
		UpdatedAt: m.UpdatedAt.Seconds,
		Content:   m.Content,
		Tags:      m.Tags,
		Sha1:      m.Sha1,
	}
}
