package services

import (
	"context"

	"github.com/9d77v/go-lib/ptrs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
)

//DeviceService ..
type DeviceService struct {
	base.Service
}

//CreateDevice  ..
func (s DeviceService) CreateDevice(ctx context.Context, in *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	resp := new(pb.CreateDeviceResponse)
	m := models.NewDeviceFromPB(in)
	deviceModel := models.NewDeviceModel()
	if err := deviceModel.GetByID(uint(in.DeviceModelId)); err != nil {
		return nil, err
	}
	m.Begin()
	err := m.Create(m)
	if err != nil {
		m.Rollback()
		return resp, err
	}
	m.AccessKey = consts.GetDeviceAccessKey(m.ID)
	m.SecretKey = consts.GetDeviceSecretKey()
	err = m.Save(m)
	if err != nil {
		m.Rollback()
		return resp, err
	}
	err = m.SaveAttributesFromAttributeModels(deviceModel.AttributeModels)
	if err != nil {
		m.Rollback()
		return resp, err
	}
	err = m.SaveTelemetriesFromTelemetryModels(deviceModel.TelemetryModels)
	if err != nil {
		m.Rollback()
		return resp, err
	}
	err = m.Commit()
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateDevice ..
func (s DeviceService) UpdateDevice(ctx context.Context, in *pb.UpdateDeviceRequest) (*pb.UpdateDeviceResponse, error) {
	resp := &pb.UpdateDeviceResponse{
		Id: in.Id,
	}
	m := models.NewDevice()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"name":     in.Name,
		"ip":       in.Ip,
		"port":     uint(in.Port),
		"username": in.Username,
		"password": in.Password,
	})
}

//ListDevice ..
func (s DeviceService) ListDevice(ctx context.Context, in *pb.ListDeviceRequest) (*pb.ListDeviceResponse, error) {
	resp := new(pb.ListDeviceResponse)
	m := models.NewDevice()
	m.FuzzyQuery(in.SearchParam.Keyword, "name")
	if in.DeviceType != nil {
		in.SearchParam.TableName = m.TableName()
		m.SelectDeviceType(ptrs.Int64(in.DeviceType))
	}
	omitFields := []string{"attributes", "telemetries",
		"deviceModel"}
	replaceFunc := func(field base.GraphQLField) error {
		if in.DeviceType != nil {
			m.SelectWithPrefix(field.Fields, m.TableName()+".", omitFields...)
		}
		if field.FieldMap["attributes"] {
			m.Preload("Attributes").Preload("Attributes.AttributeModel")
		}
		if field.FieldMap["telemetries"] {
			m.Preload("Telemetries").Preload("Telemetries.TelemetryModel")
		}
		if field.FieldMap["deviceModel"] {
			m.Preload("DeviceModel")
		}
		return nil
	}
	data := make([]*models.Device, 0)
	total, err := s.GetNewConnection(m, in.SearchParam, &data, replaceFunc, omitFields...)
	resp.TotalCount = total
	resp.Edges = m.ToDevicePBs(data)
	return resp, err
}
