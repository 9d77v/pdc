package models

import (
	"reflect"
	"sort"
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
)

//Note ...
type Note struct {
	base.Model
	ID        string         `gorm:"column:id;size:36;NOT NULL;PRIMARY_KEY"`
	ParentID  string         `gorm:"column:parent_id;size:36;NOT NULL;index:parent_id"`
	UID       int64          `gorm:"NOT NULL;index:uid"`
	NoteType  int8           `gorm:"NOT NULL"`
	Level     int8           `gorm:"NOT NULL"`
	Title     string         `gorm:"size:50;NOT NULL;"`
	State     int8           `gorm:"NOT NULL"`
	Version   int32          `gorm:"NOT NULL"`
	Color     string         `gorm:"size:7;NOT NULL;default:''"`
	Content   string         `gorm:"NOT NULL"`
	Tags      pq.StringArray `gorm:"type:varchar(10)[]"`
	SHA1      string         `gorm:"column:sha1;size:40;NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewNote ..
func NewNote() *Note {
	m := new(Note)
	m.SetDB(db.GetDB())
	return m
}

//NewNoteFromPB ...
func NewNoteFromPB(in *pb.Note) *Note {
	now := time.Now()
	createdAt := now
	updatedAt := now
	if in.CreatedAt != nil {
		createdAt, _ = ptypes.Timestamp(in.CreatedAt)
	}
	if in.UpdatedAt != nil {
		updatedAt, _ = ptypes.Timestamp(in.UpdatedAt)
	}
	note := &Note{
		ID:        in.Id,
		ParentID:  in.ParentId,
		UID:       in.Uid,
		NoteType:  int8(in.NoteType),
		Level:     int8(in.Level),
		Title:     in.Title,
		Version:   in.Version,
		Color:     in.Color,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Content:   in.Content,
		Tags:      in.Tags,
		SHA1:      in.Sha1,
	}
	return note
}

//TableName ..
func (m *Note) TableName() string {
	return db.TablePrefix() + "note"
}

//HasInvalidNotes check if uid inconsistent or level error
func (m *Note) HasInvalidNotes(in *pb.SyncNotesRequest) bool {
	for _, v := range in.UnsyncedNotes {
		if v.Uid != in.Uid {
			return true
		}
		switch v.NoteType {
		case pb.NoteType_Directory:
			if v.Level != 1 && v.Level != 2 {
				return true
			}
		default:
			if v.Level != 3 {
				return true
			}
		}
	}
	return false
}

//GetNotesForSync ...
func (m *Note) GetNotesForSync(in *pb.SyncNotesRequest) ([]*Note, error) {
	updateNoteIDs := make([]string, 0, len(in.UnsyncedNotes))
	for _, v := range in.UnsyncedNotes {
		updateNoteIDs = append(updateNoteIDs, v.Id)
	}
	lastUpdateTime := time.Unix(0, 0)
	if in.LastUpdateTime != nil {
		lastUpdateTime, _ = ptypes.Timestamp(in.LastUpdateTime)
	}
	list := make([]*Note, 0, len(in.UnsyncedNotes))
	m.IDQuery(in.Uid, "uid")
	if len(updateNoteIDs) > 0 {
		m.Where("(updated_at > ?) or id in(?)", lastUpdateTime, updateNoteIDs)
	} else {
		m.Where("updated_at > ?", lastUpdateTime)
	}
	err := m.Find(&list)
	return list, err
}

//Sorts ..
func (m *Note) Sorts(notes []*Note) []*Note {
	sort.SliceStable(notes, func(i, j int) bool {
		return notes[i].UpdatedAt.Unix() < notes[j].UpdatedAt.Unix() && notes[i].ID < notes[j].ID
	})
	return notes
}

//ListToMap list 转map
func (m *Note) ListToMap(notes []*Note) map[string]*Note {
	noteMap := make(map[string]*Note)
	for _, v := range notes {
		noteMap[v.ID] = v
	}
	return noteMap
}

//ClassifyNotes classify notes for create,update,delete and client should sync
func (m *Note) ClassifyNotes(serverNotes []*Note,
	clientNotes []*pb.Note) ([]*Note, []*Note, []*Note, []*Note) {
	serverNoteMap := m.ListToMap(serverNotes)
	createNotes := make([]*Note, 0)
	updateNotes := make([]*Note, 0)
	deleteNotes := make([]*Note, 0)
	clientShouldSyncNotes := make([]*Note, 0)
	for _, clientNote := range clientNotes {
		currentNote := serverNoteMap[clientNote.Id]
		if currentNote == nil {
			clientNote.CreatedAt = nil
			clientNote.UpdatedAt = nil
			createNotes = append(createNotes, NewNoteFromPB(clientNote))
		} else if clientNote.State == pb.NoteState_IsDeleted {
			if currentNote.State == int8(pb.NoteState_IsDeleted) {
				clientShouldSyncNotes = append(clientShouldSyncNotes, currentNote)
			} else {
				deleteNotes = append(deleteNotes, NewNoteFromPB(clientNote))
			}
		} else if currentNote.canUpdate(clientNote) {
			clientNote.Version = currentNote.Version + 1
			updateNotes = append(updateNotes, NewNoteFromPB(clientNote))
		} else {
			clientShouldSyncNotes = append(clientShouldSyncNotes, currentNote)
		}
		delete(serverNoteMap, clientNote.Id)
	}
	for _, v := range serverNoteMap {
		clientShouldSyncNotes = append(clientShouldSyncNotes, v)
	}
	return createNotes, updateNotes, deleteNotes, clientShouldSyncNotes
}

func (m *Note) canUpdate(in *pb.Note) bool {
	if m.Version > in.Version {
		return false
	}
	if m.Title != in.Title || m.Color != in.Color || m.ParentID != in.ParentId ||
		!reflect.DeepEqual(([]string)(m.Tags), in.Tags) || m.State != int8(in.State) {
		return true
	}
	switch in.NoteType {
	case pb.NoteType_File:
		if m.SHA1 != in.Sha1 {
			return true
		}
	}
	return false
}

//UpdateNotes ...
func (m *Note) UpdateNotes(notes []*Note) error {
	now := time.Now()
	for _, v := range notes {
		err := m.Table(m.TableName()).IDQuery(v.ID).Updates(map[string]interface{}{
			"parent_id":  v.ParentID,
			"level":      v.Level,
			"title":      v.Title,
			"state":      v.State,
			"version":    v.Version,
			"color":      v.Color,
			"content":    v.Content,
			"tags":       v.Tags,
			"sha1":       v.SHA1,
			"updated_at": now,
		})
		if err != nil {
			return err
		}
		if v.NoteType == int8(pb.NoteType_File) {
			err = m.Create(NewNoteHistory(v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//GetLatestNotes ..
func (m *Note) GetLatestNotes(notes []*Note) []*Note {
	size := len(notes)
	result := make([]*Note, 0, size)
	ids := make([]string, 0, size)
	for _, v := range notes {
		ids = append(ids, v.ID)
	}
	if len(ids) > 0 {
		m.IDArrayQuery(ids).Find(&result)
	}
	return result
}

//DeleteAllChildNodes ...
func (m *Note) DeleteAllChildNodes() (err error) {
	err = m.IDQuery(m.ID).Table(m.TableName()).Updates(map[string]interface{}{
		"state":      int8(pb.NoteState_IsDeleted),
		"updated_at": time.Now(),
	})
	if m.NoteType == int8(pb.NoteType_Directory) {
		notes, subErr := m.List(m.ID, m.UID)
		if err != nil {
			return subErr
		}
		for _, note := range notes {
			note.SetDB(m.GetDB())
			err = note.DeleteAllChildNodes()
			if err != nil {
				return err
			}
		}
	}
	return err
}

//List list model
func (m *Note) List(parentID string, uid int64) (list []*Note, err error) {
	list = make([]*Note, 0)
	m.Where("parent_id=? and uid=?", parentID, uid)
	err = m.Find(&list)
	return list, err
}

//ToNotePBs ...
func (m *Note) ToNotePBs(data []*Note) []*pb.Note {
	result := make([]*pb.Note, 0, len(data))
	for _, v := range data {
		result = append(result, v.toNotePB())
	}
	return result
}

//toNotePB ...
func (m *Note) toNotePB() *pb.Note {
	createdAt, _ := ptypes.TimestampProto(m.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(m.UpdatedAt)
	return &pb.Note{
		Id:        m.ID,
		ParentId:  m.ParentID,
		Uid:       m.UID,
		NoteType:  pb.NoteType(m.NoteType),
		Level:     int32(m.Level),
		Title:     m.Title,
		State:     pb.NoteState(m.State),
		Version:   m.Version,
		Color:     m.Color,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Content:   m.Content,
		Tags:      m.Tags,
		Sha1:      m.SHA1,
	}
}
