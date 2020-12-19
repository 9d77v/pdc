package services

import (
	"github.com/9d77v/pdc/internal/module/device-service/models"
)

func getDefaultHikvisionAttributeModels(id uint) []*models.AttributeModel {
	return []*models.AttributeModel{
		{
			DeviceModelID: id,
			Key:           "device_name",
			Name:          "设备名称",
		},
		{
			DeviceModelID: id,
			Key:           "device_id",
			Name:          "设备ID",
		},
		{
			DeviceModelID: id,
			Key:           "device_description",
			Name:          "设备描述",
		},
		{
			DeviceModelID: id,
			Key:           "device_location",
			Name:          "设备位置",
		},
		{
			DeviceModelID: id,
			Key:           "system_contact",
			Name:          "系统联系方",
		},
		{
			DeviceModelID: id,
			Key:           "model",
			Name:          "类型",
		},
		{
			DeviceModelID: id,
			Key:           "serial_number",
			Name:          "序列号",
		},
		{
			DeviceModelID: id,
			Key:           "mac_address",
			Name:          "MAC地址",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_version",
			Name:          "固件版本",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_releasedDate",
			Name:          "固件发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "encoder_version",
			Name:          "编码器版本",
		},
		{
			DeviceModelID: id,
			Key:           "encoder_released_date",
			Name:          "编码器发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "boot_version",
			Name:          "引导版本",
		},
		{
			DeviceModelID: id,
			Key:           "boot_released_date",
			Name:          "引导发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "hardware_version",
			Name:          "硬件版本",
		},
		{
			DeviceModelID: id,
			Key:           "device_type",
			Name:          "设备类型",
		},
		{
			DeviceModelID: id,
			Key:           "telecontrol_id",
			Name:          "远程 ID",
		},
		{
			DeviceModelID: id,
			Key:           "support beep",
			Name:          "支持蜂鸣音",
		},
		{
			DeviceModelID: id,
			Key:           "support_video_loss",
			Name:          "支持视频丢失",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_version_info",
			Name:          "固件版本信息",
		},
	}
}
