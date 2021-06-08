package services

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/note-service/models"
	"github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/golang/protobuf/ptypes"
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
		&models.NoteHistory{},
		&models.Note{},
	}
}

func initDB() {
	config := &db.DBConfig{
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

func TestNoteService_SyncNotes(t *testing.T) {
	tnow := time.Unix(111, 0)
	now, _ := ptypes.TimestampProto(tnow)
	serverNotes := []*models.Note{
		{
			ID:  "3",
			UID: 1,
		},
		{
			ID:  "4",
			UID: 1,
		},
		{
			ID:        "5",
			UID:       1,
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "6",
			UID:       1,
			State:     2,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "7",
			UID:       1,
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "8",
			UID:       1,
			Version:   10,
			State:     1,
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
		{
			ID:        "9",
			UID:       1,
			NoteType:  1,
			SHA1:      "33",
			CreatedAt: tnow,
			UpdatedAt: tnow,
		},
	}
	models.NewNote().Create(&serverNotes)
	type args struct {
		ctx context.Context
		in  *pb.SyncNotesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test note format error ", args{context.Background(), &pb.SyncNotesRequest{
				Uid: 1,
				UnsyncedNotes: []*pb.Note{
					{
						Id:        "2",
						Uid:       2,
						NoteType:  pb.NoteType_Directory,
						Level:     1,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}}}, true,
		},
		{"test syncnotes ok ", args{context.Background(), &pb.SyncNotesRequest{
			Uid: 1,
			UnsyncedNotes: []*pb.Note{
				{
					Id:        "2",
					Uid:       1,
					NoteType:  pb.NoteType_Directory,
					Level:     1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        "5",
					Uid:       1,
					NoteType:  pb.NoteType_File,
					Level:     3,
					State:     pb.NoteState_IsDeleted,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        "6",
					Uid:       1,
					NoteType:  pb.NoteType_Directory,
					Level:     2,
					State:     pb.NoteState_IsDeleted,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        "7",
					Uid:       1,
					NoteType:  pb.NoteType_File,
					Level:     3,
					Title:     "dd",
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        "8",
					Uid:       1,
					Version:   1,
					NoteType:  pb.NoteType_File,
					Level:     3,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        "9",
					Uid:       1,
					NoteType:  pb.NoteType_File,
					Level:     3,
					Sha1:      "22",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NoteService{}.SyncNotes(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteService.SyncNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
