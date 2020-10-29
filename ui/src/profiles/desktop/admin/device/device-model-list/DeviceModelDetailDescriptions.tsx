import { Descriptions } from 'antd';
import React from 'react'
import { IDeviceModel } from 'src/consts/consts';
import { formatDetailTime } from 'src/utils/util';

export const DeviceModelDetailDescriptions: React.FC<IDeviceModel> = ({
    id,
    name,
    desc,
    deviceType,
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
            <Descriptions.Item label="描述">{desc}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{formatDetailTime(createdAt)}</Descriptions.Item>
            <Descriptions.Item label="更新时间">{formatDetailTime(updatedAt)}</Descriptions.Item>
        </Descriptions>
    );
};