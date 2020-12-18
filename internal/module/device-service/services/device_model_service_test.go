package services

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/stretchr/testify/assert"
)

const testDBName = "pdc_test"

func TestMain(m *testing.M) {
	initDB()
	m.Run()
	clean()
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
	err := db.GetDB(config).AutoMigrate(
		&models.DeviceModel{},
		&models.TelemetryModel{},
		&models.AttributeModel{},
		&models.Device{},
		&models.Attribute{},
		&models.Telemetry{},
		&models.DeviceDashboard{},
		&models.DeviceDashboardTelemetry{},
		&models.DeviceDashboardCamera{},
		&models.CameraTimeLapseVideo{},
	)
	if err != nil {
		fmt.Println("auto migrate failed:", err)
	}
}

func clean() {
	err := db.GetDB().Where("1 = 1").Unscoped().Delete(&models.DeviceDashboardTelemetry{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.DeviceDashboardCamera{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.DeviceDashboard{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.Attribute{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.Telemetry{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.Device{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.AttributeModel{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.TelemetryModel{}).Error
	checkErr(err)
	err = db.GetDB().Where("1 = 1").Unscoped().Delete(&models.DeviceModel{}).Error
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}

var (
	testDeviceModel = &pb.CreateDeviceModelRequest{
		Name:          "测试模型1",
		Desc:          "desc",
		DeviceType:    pb.DeviceType_Camera,
		CameraCompany: pb.CameraCompany_Hikvision,
	}
	deviceModelService = DeviceModelService{}
	ctx                = context.Background()
)

func TestDeviceModelService_CreateDeviceModel(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.CreateDeviceModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateDeviceModel", args{ctx, testDeviceModel}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.CreateDeviceModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_UpdateDeviceModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	type args struct {
		ctx context.Context
		in  *pb.UpdateDeviceModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test deviceModel exist", args{ctx, &pb.UpdateDeviceModelRequest{
			Id:   deviceModel.Id,
			Name: "测试模型2",
			Desc: "desc2",
		}}, false},
		{"test deviceModel not exist", args{ctx, &pb.UpdateDeviceModelRequest{
			Id:   deviceModel.Id + 1,
			Name: "测试模型2",
			Desc: "desc2",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.UpdateDeviceModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_CreateAttributeModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	attributeModel := &pb.CreateAttributeModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateAttributeModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateAttributeModel", args{ctx, attributeModel}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.CreateAttributeModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_UpdateAttributeModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testAttributeModel := &pb.CreateAttributeModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
	}
	attributeModel, _ := deviceModelService.CreateAttributeModel(ctx, testAttributeModel)
	type args struct {
		ctx context.Context
		in  *pb.UpdateAttributeModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test attributeModel exist", args{ctx, &pb.UpdateAttributeModelRequest{
			Id:   attributeModel.Id,
			Name: "name2",
		}}, false},
		{"test attributeModel not exist", args{ctx, &pb.UpdateAttributeModelRequest{
			Id:   attributeModel.Id + 1,
			Name: "name2",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.UpdateAttributeModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_DeleteAttributeModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testAttributeModel := &pb.CreateAttributeModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
	}
	attributeModel, _ := deviceModelService.CreateAttributeModel(ctx, testAttributeModel)
	type args struct {
		ctx context.Context
		in  *pb.DeleteAttributeModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ids length is zero", args{ctx, &pb.DeleteAttributeModelRequest{}}, false},
		{"test ids length is  not zero and record exist", args{ctx, &pb.DeleteAttributeModelRequest{
			Ids: []int64{int64(attributeModel.Id)},
		}}, false},
		{"test ids length is  not zero and record not exist", args{ctx, &pb.DeleteAttributeModelRequest{
			Ids: []int64{int64(attributeModel.Id + 1)},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := deviceModelService.DeleteAttributeModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeviceModelService_CreateTelemetryModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	telemetryModel := &pb.CreateTelemetryModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
		Factor:        1,
		Unit:          "A",
		UnitName:      "安培",
		Scale:         2,
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateTelemetryModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateTelemetryModel", args{ctx, telemetryModel}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.CreateTelemetryModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_UpdateTelemetryModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testTelemetryModel := &pb.CreateTelemetryModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
		Factor:        1,
		Unit:          "A",
		UnitName:      "安培",
		Scale:         2,
	}
	telemetryModel, _ := deviceModelService.CreateTelemetryModel(ctx, testTelemetryModel)
	type args struct {
		ctx context.Context
		in  *pb.UpdateTelemetryModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test telemetryModel exist", args{ctx, &pb.UpdateTelemetryModelRequest{
			Id:   telemetryModel.Id,
			Name: "name2",
		}}, false},
		{"test telemetryModel not exist", args{ctx, &pb.UpdateTelemetryModelRequest{
			Id:   telemetryModel.Id + 1,
			Name: "name2",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceModelService.UpdateTelemetryModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceModelService_DeleteTelemetryModel(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testTelemetryModel := &pb.CreateTelemetryModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
		Factor:        1,
		Unit:          "A",
		UnitName:      "安培",
		Scale:         2,
	}
	telemetryModel, _ := deviceModelService.CreateTelemetryModel(ctx, testTelemetryModel)
	type args struct {
		ctx context.Context
		in  *pb.DeleteTelemetryModelRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ids length is zero", args{ctx, &pb.DeleteTelemetryModelRequest{}}, false},
		{"test ids length is  not zero and record exist", args{ctx, &pb.DeleteTelemetryModelRequest{
			Ids: []int64{int64(telemetryModel.Id)},
		}}, false},
		{"test ids length is  not zero and record not exist", args{ctx, &pb.DeleteTelemetryModelRequest{
			Ids: []int64{int64(telemetryModel.Id + 1)},
		}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := deviceModelService.DeleteTelemetryModel(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
