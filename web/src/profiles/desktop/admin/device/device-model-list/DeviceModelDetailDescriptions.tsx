import { Descriptions } from 'antd';
import React, { FC } from 'react'
import { CameraCompanyMap } from 'src/consts/consts';
import { IDeviceModel } from 'src/models/device';
import { formatDetailTime } from 'src/utils/util';

export const DeviceModelDetailDescriptions: FC<IDeviceModel> = ({
    id,
    name,
    desc,
    deviceType,
    cameraCompany,
    createdAt,
    updatedAt
}) => {
    console.log(cameraCompany)
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
            <Descriptions.Item label="描述">{desc}</Descriptions.Item>
            {deviceType === 1 ? <Descriptions.Item label="摄像头厂家">{CameraCompanyMap.get(cameraCompany)}</Descriptions.Item> : null}
            <Descriptions.Item label="创建时间">{formatDetailTime(createdAt)}</Descriptions.Item>
            <Descriptions.Item label="更新时间">{formatDetailTime(updatedAt)}</Descriptions.Item>
        </Descriptions>
    );
};