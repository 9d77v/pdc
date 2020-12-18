package services

import (
	"context"

	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

//DeviceModelService ..
type DeviceModelService struct {
	base.Service
}

//CreateDeviceModel  ..
func (s DeviceModelService) CreateDeviceModel(ctx context.Context,
	in *pb.CreateDeviceModelRequest) (*pb.CreateDeviceModelResponse, error) {
	resp := new(pb.CreateDeviceModelResponse)
	m := models.NewDeviceModelFromPB(in)
	err := m.Create(m)
	if err != nil {
		return resp, err
	}
	if m.IsHikivisionCamera(in.DeviceType, in.CameraCompany) {
		attributes := getDefaultHikvisionAttributeModels(m.ID)
		err = m.Create(&attributes)
	}
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateDeviceModel ..
func (s DeviceModelService) UpdateDeviceModel(ctx context.Context,
	in *pb.UpdateDeviceModelRequest) (*pb.UpdateDeviceModelResponse, error) {
	resp := &pb.UpdateDeviceModelResponse{
		Id: in.Id,
	}
	m := models.NewDeviceModel()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"name": in.Name,
		"desc": in.Desc,
	})
}

//CreateAttributeModel  ..
func (s DeviceModelService) CreateAttributeModel(ctx context.Context,
	in *pb.CreateAttributeModelRequest) (*pb.CreateAttributeModelResponse, error) {
	resp := new(pb.CreateAttributeModelResponse)
	m := models.NewAttributeModelFromPB(in)
	err := m.Create(m)
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateAttributeModel ..
func (s DeviceModelService) UpdateAttributeModel(ctx context.Context,
	in *pb.UpdateAttributeModelRequest) (*pb.UpdateAttributeModelResponse, error) {
	resp := &pb.UpdateAttributeModelResponse{
		Id: in.Id,
	}
	m := models.NewAttributeModel()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"name": in.Name,
	})
}

//DeleteAttributeModel ..
func (s DeviceModelService) DeleteAttributeModel(ctx context.Context,
	in *pb.DeleteAttributeModelRequest) (*pb.DeleteAttributeModelResponse, error) {
	resp := &pb.DeleteAttributeModelResponse{
		Ids: in.Ids,
	}
	m := models.NewAttributeModel()
	return resp, m.Delete(m, in.Ids)
}

//CreateTelemetryModel  ..
func (s DeviceModelService) CreateTelemetryModel(ctx context.Context,
	in *pb.CreateTelemetryModelRequest) (*pb.CreateTelemetryModelResponse, error) {
	resp := new(pb.CreateTelemetryModelResponse)
	m := models.NewTelemetryModelFromPB(in)
	err := m.Create(m)
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateTelemetryModel ..
func (s DeviceModelService) UpdateTelemetryModel(ctx context.Context,
	in *pb.UpdateTelemetryModelRequest) (*pb.UpdateTelemetryModelResponse, error) {
	resp := &pb.UpdateTelemetryModelResponse{
		Id: in.Id,
	}
	m := models.NewTelemetryModel()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"name":      in.Name,
		"factor":    in.Factor,
		"unit":      in.Unit,
		"unit_name": in.UnitName,
		"scale":     in.Scale,
	})
}

//DeleteTelemetryModel ..
func (s DeviceModelService) DeleteTelemetryModel(ctx context.Context,
	in *pb.DeleteTelemetryModelRequest) (*pb.DeleteAttributeModelResponse, error) {
	resp := &pb.DeleteAttributeModelResponse{
		Ids: in.Ids,
	}
	m := models.NewTelemetryModel()
	return resp, m.Delete(m, in.Ids)
}

//ListDeviceModel ..
func (s DeviceModelService) ListDeviceModel(ctx context.Context, searchParam model.SearchParam) (int64, []*model.DeviceModel, error) {
	var deviceModel base.Repository = models.NewDeviceModel()
	deviceModel.FuzzyQuery(searchParam.Keyword, "name")
	replaceFunc := func(field base.GraphQLField) error {
		if field.FieldMap["attributeModels"] {
			deviceModel.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.AttributeModel{})
			})
		}
		if field.FieldMap["telemetryModels"] {
			deviceModel.Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.TelemetryModel{})
			})
		}
		return nil
	}
	data := make([]*models.DeviceModel, 0)
	total, err := s.GetConnection(ctx, deviceModel, searchParam, &data, replaceFunc,
		"attributeModels", "telemetryModels")
	return total, getDeviceModels(data), err
}
