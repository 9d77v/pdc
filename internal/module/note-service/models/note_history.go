package models

import (
	"time"

	"github.com/9d77v/pdc/internal/module/base"
)

//NoteHistory ...
type NoteHistory struct {
	base.Model
	ID        int64     `gorm:"column:id;size:36;NOT NULL;PRIMARY_KEY"`
	NoteID    string    `gorm:"column:note_id;size:36;NOT NULL;index:note_id"`
	UID       int64     `gorm:"NOT NULL;index:uid"`
	Title     string    `gorm:"size:50;NOT NULL"`
	Content   string    `gorm:"NOT NULL"`
	Version   int32     `gorm:"NOT NULL"`
	SHA1      string    `gorm:"column:sha1;size:40;NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL"`
}

//NewNoteHistory ...
func NewNoteHistory(in *Note) *NoteHistory {
	return &NoteHistory{
		NoteID:  in.ID,
		UID:     in.UID,
		Title:   in.Title,
		Content: in.Content,
		SHA1:    in.SHA1,
		Version: in.Version,
	}
}
