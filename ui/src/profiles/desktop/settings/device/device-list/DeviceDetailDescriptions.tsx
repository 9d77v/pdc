import { Descriptions } from 'antd';
import React from 'react'
import { IDevice } from '../../../../../consts/consts';
import { formatDetailTime } from '../../../../../utils/util';

export const DeviceDetailDescriptions: React.FC<IDevice> = ({
    id,
    name,
    deviceModelName,
    deviceModelDesc,
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
            <Descriptions.Item label="设备模板名称">{deviceModelName}</Descriptions.Item>
            <Descriptions.Item label="设备模板描述">{deviceModelDesc}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{formatDetailTime(createdAt)}</Descriptions.Item>
            <Descriptions.Item label="更新时间">{formatDetailTime(updatedAt)}</Descriptions.Item>
        </Descriptions>
    )
}