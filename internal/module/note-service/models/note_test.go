package models

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

const testDBName = "pdc_test"

func TestMain(m *testing.M) {
	initDB()
	clean()
	m.Run()
	clean()
}

func getStructs() []interface{} {
	return []interface{}{
		&NoteHistory{},
		&Note{},
	}
}

func initDB() {
	config := &config.DBConfig{
		Driver:       "postgres",
		Host:         "domain.local",
		Port:         5432,
		User:         "postgres",
		Password:     "123456",
		Name:         testDBName,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    true,
	}
	err := db.GetDB(config).AutoMigrate(getStructs()...)
	if err != nil {
		fmt.Println("auto migrate failed:", err)
	}
}

func clean() {
	for _, v := range getStructs() {
		err := db.GetDB().Where("1 = 1").Unscoped().Delete(v).Error
		if err != nil {
			log.Println("error:", err)
		}
	}
}

var (
	tnow      = time.Unix(111, 0)
	now, _    = ptypes.TimestampProto(tnow)
	testNotes = []*Note{
		{
			UID:       1,
			ID:        "1",
			NoteType:  0,
			Level:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			UID:       1,
			ID:        "2",
			ParentID:  "1",
			NoteType:  0,
			Level:     2,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			UID:       1,
			ID:        "3",
			ParentID:  "2",
			NoteType:  1,
			Level:     3,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
	}
)

func TestNote_HasInvalidNotes(t *testing.T) {
	type args struct {
		in *pb.SyncNotesRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test uid inconsistent", args{
			&pb.SyncNotesRequest{
				Uid: 1,
				UnsyncedNotes: []*pb.Note{
					{
						Uid: 2,
					},
				},
			},
		}, true},
		{"test directory level error", args{
			&pb.SyncNotesRequest{
				Uid: 1,
				UnsyncedNotes: []*pb.Note{
					{
						Uid:      1,
						Id:       "12345",
						NoteType: pb.NoteType_Directory,
						Level:    4,
					},
				},
			},
		}, true},
		{"test file level error", args{
			&pb.SyncNotesRequest{
				Uid: 1,
				UnsyncedNotes: []*pb.Note{
					{
						Uid:      1,
						Id:       "123456",
						NoteType: pb.NoteType_File,
						Level:    1,
					},
				},
			},
		}, true},
		{"test ok", args{
			&pb.SyncNotesRequest{
				Uid: 1,
				UnsyncedNotes: []*pb.Note{
					{
						Uid:      1,
						Id:       "12345",
						NoteType: pb.NoteType_Directory,
						Level:    1,
					},
					{
						Uid:      1,
						Id:       "123456",
						NoteType: pb.NoteType_File,
						Level:    3,
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNote().HasInvalidNotes(tt.args.in); got != tt.want {
				t.Errorf("Note.HasInvalidNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_GetNotesForSync(t *testing.T) {
	lastUpdateTime := time.Unix(100, 0)
	lastUpdateTimeStamp, _ := ptypes.TimestampProto(lastUpdateTime)
	oldTnow := time.Unix(11, 0)
	oldNow, _ := ptypes.TimestampProto(oldTnow)
	additionalInputNotes := []*Note{
		{
			UID:       1,
			ID:        "01",
			NoteType:  0,
			Level:     1,
			CreatedAt: oldTnow,
			UpdatedAt: oldTnow,
		},
		{
			UID:       1,
			ID:        "02",
			NoteType:  0,
			Level:     1,
			CreatedAt: oldTnow,
			UpdatedAt: oldTnow,
		},
	}
	NewNote().Create(append(testNotes, additionalInputNotes...))
	additionalGotNotes := []*Note{
		{
			UID:       1,
			ID:        "02",
			NoteType:  0,
			Level:     1,
			CreatedAt: oldTnow,
			UpdatedAt: oldTnow,
		},
	}
	type args struct {
		in *pb.SyncNotesRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*Note
		wantErr bool
	}{
		{"test has UnsyncedNotes", args{
			&pb.SyncNotesRequest{
				Uid:            1,
				LastUpdateTime: lastUpdateTimeStamp,
				UnsyncedNotes: []*pb.Note{
					{Uid: 1,
						Id:        "02",
						NoteType:  pb.NoteType_Directory,
						Level:     1,
						CreatedAt: oldNow,
						UpdatedAt: oldNow,
						Title:     "ttt",
					},
				},
			},
		}, append(testNotes, additionalGotNotes...), false},
		{"test has not UnsyncedNotes", args{
			&pb.SyncNotesRequest{
				Uid:            1,
				LastUpdateTime: lastUpdateTimeStamp,
				UnsyncedNotes:  []*pb.Note{},
			},
		}, append(testNotes), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNote().GetNotesForSync(tt.args.in)
			assert.False(t, (err != nil) != tt.wantErr)
			assert.Equal(t, NewNote().Sorts(got), NewNote().Sorts(tt.want))
		})
	}
}

func TestNote_ClassifyNotes(t *testing.T) {
	createClientNote := &pb.Note{
		Id:        "2",
		CreatedAt: now,
		UpdatedAt: now,
	}
	updateClientNote := []*pb.Note{
		{
			Id:        "7",
			Title:     "dd",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Id:        "8",
			Version:   1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Id:        "9",
			NoteType:  pb.NoteType_File,
			Sha1:      "22",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	deleteClientNote := []*pb.Note{
		{
			Id:        "5",
			State:     pb.NoteState_IsDeleted,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Id:        "6",
			State:     pb.NoteState_IsDeleted,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	oldServerNotes := []*Note{
		{
			ID: "3",
		},
		{
			ID: "4",
		},
	}
	updateServerNotes := []*Note{
		{
			ID:        "7",
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "8",
			Version:   10,
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "9",
			NoteType:  1,
			SHA1:      "33",
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
	}
	deleteServerNotes := []*Note{
		{
			ID:        "5",
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "6",
			State:     2,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
	}
	gotUpdateClientNotes := NewNoteFromPB(updateClientNote[0])
	gotUpdateClientNotes.Version++
	gotUpdateClientNotes2 := NewNoteFromPB(updateClientNote[2])
	gotUpdateClientNotes2.Version++
	type args struct {
		serverNotes []*Note
		clientNotes []*pb.Note
	}
	tests := []struct {
		name  string
		args  args
		want  []*Note
		want1 []*Note
		want2 []*Note
		want3 []*pb.Note
	}{
		{"test classifyNotes", args{
			append(append(append([]*Note{
				{
					ID: "1",
				},
			}, oldServerNotes...), updateServerNotes...), deleteServerNotes...),
			append(append([]*pb.Note{
				{
					Id: "1",
				},
				createClientNote}, updateClientNote...), deleteClientNote...),
		}, []*Note{
			NewNoteFromPB(createClientNote),
		}, []*Note{
			gotUpdateClientNotes,
			gotUpdateClientNotes2,
		}, []*Note{
			NewNoteFromPB(deleteClientNote[0]),
		},
			append([]*pb.Note{
				{
					Id: "1",
				},
				updateClientNote[1],
				deleteClientNote[1],
			}, NewNote().ToNotePBs(oldServerNotes)...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := NewNote().ClassifyNotes(tt.args.serverNotes, tt.args.clientNotes)
			assert.Equal(t, len(got), len(tt.want))
			assert.Equal(t, got1, tt.want1)
			assert.Equal(t, got2, tt.want2)
			assert.Equal(t, got3, tt.want3)
		})
	}
}

func TestNote_UpdateNotes(t *testing.T) {
	NewNote().Create(testNotes)
	updateNotes := []*Note{
		{
			UID:       1,
			ID:        "1",
			NoteType:  0,
			Level:     1,
			Title:     "www",
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			UID:       1,
			ID:        "2",
			ParentID:  "1",
			Title:     "sss",
			NoteType:  0,
			Level:     2,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			UID:       1,
			ID:        "3",
			Title:     "vvv",
			ParentID:  "2",
			NoteType:  1,
			Level:     3,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
	}
	type args struct {
		in []*Note
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test update notes", args{updateNotes}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewNote().UpdateNotes(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Note.UpdateNotes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	clean()
}

func TestNote_GetLatestNotes(t *testing.T) {
	NewNote().Create(testNotes)
	type args struct {
		notes []*Note
	}
	tests := []struct {
		name string
		args args
		want []*Note
	}{
		{"test GetLatestNotes", args{testNotes}, testNotes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNote().GetLatestNotes(tt.args.notes); !assert.Equal(t, got, tt.want) {
				t.Errorf("Note.GetLatestNotes() = %v, want %v", got, tt.want)
			}
		})
	}
	clean()
}

func TestNote_DeleteAllChildNodes(t *testing.T) {
	NewNote().Create(testNotes)
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test DeleteAllChildNodes", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewNote()
			m.ID = "1"
			m.UID = 1
			if err := m.DeleteAllChildNodes(); (err != nil) != tt.wantErr {
				t.Errorf("Note.DeleteAllChildNodes() error = %v, wantErr %v", err, tt.wantErr)
			}
			notes := make([]*Note, 0)
			m.Where("state!=?", pb.NoteState_IsDeleted).Find(&notes)
			assert.Zero(t, len(notes))
		})
	}
}
