import { Descriptions } from 'antd';
import React, { FC } from 'react'
import { CameraCompanyMap } from 'src/consts/consts';
import { IDevice } from 'src/models/device';
import { formatDetailTime } from 'src/utils/util';

export const DeviceDetailDescriptions: FC<IDevice> = ({
    id,
    name,
    ip,
    port,
    accessKey,
    secretKey,
    username,
    password,
    deviceModelName,
    deviceModelDesc,
    deviceModelDeviceType,
    deviceModelCameraCompany,
    createdAt,
    updatedAt
}) => {
    return (
        <Descriptions
            title="设备详情"
            bordered column={1}
            style={{
                textAlign: "left",
                padding: "0px 10px 10px 10px",
            }}>
            <Descriptions.Item label="ID">{id}</Descriptions.Item>
            <Descriptions.Item label="名称">{name}</Descriptions.Item>
            <Descriptions.Item label="IP">{ip}</Descriptions.Item>
            <Descriptions.Item label="端口">{port}</Descriptions.Item>
            <Descriptions.Item label="AccessKey">{accessKey}</Descriptions.Item>
            <Descriptions.Item label="SecretKey">{secretKey}</Descriptions.Item>
            <Descriptions.Item label="用户名">{username}</Descriptions.Item>
            <Descriptions.Item label="密码">{password}</Descriptions.Item>
            <Descriptions.Item label="设备模板名称">{deviceModelName}</Descriptions.Item>
            <Descriptions.Item label="设备模板描述">{deviceModelDesc}</Descriptions.Item>
            {deviceModelDeviceType === 1 ? <Descriptions.Item label="摄像头厂家">{CameraCompanyMap.get(deviceModelCameraCompany)}</Descriptions.Item> : null}
            <Descriptions.Item label="创建时间">{formatDetailTime(createdAt)}</Descriptions.Item>
            <Descriptions.Item label="更新时间">{formatDetailTime(updatedAt)}</Descriptions.Item>
        </Descriptions>
    )
}
