import { List, Button, Tag } from 'antd';
import React, { useState } from 'react';
import { DeviceModelCreateForm, INewDeviceModel } from './DeviceModelCreateForm';
import { ADD_DEVICE_MODEL, LIST_DEVICE_MODEL } from '../../../../../consts/device.gql';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { DeviceTypeMap, IDeviceModel } from '../../../../../consts/consts';
import "../../../../../style/card.less"

interface IDeviceModelListProps {
    currentSelectID: number
    setCurrentSelectItem: (item: IDeviceModel) => void
}

export const DeviceModelList = (props: IDeviceModelListProps) => {
    const [deviceModelCreateFormVisible, setDeviceModelCreateFormVisible] = useState(false)
    const [addDeviceModel] = useMutation(ADD_DEVICE_MODEL);
    const [pagination, setPagination] = useState({
        current: 1,
    })
    const { loading, data, refetch, } = useQuery(LIST_DEVICE_MODEL,
        {
            variables: {
                page: pagination.current,
                pageSize: 7,
                sorts: [{
                    field: 'id',
                    isAsc: false
                }]
            },
            fetchPolicy: "cache-and-network"
        });

    const onDeviceTypeCreate = async (values: INewDeviceModel) => {
        await addDeviceModel({
            variables: {
                "input": {
                    "name": values.name,
                    "deviceType": values.deviceType || 0,
                    "desc": values.desc,
                }
            }
        });
        setDeviceModelCreateFormVisible(false);
        await refetch()
    };
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setDeviceModelCreateFormVisible(true)
                }}
                style={{
                    float: 'left',
                    marginLeft: 20,
                    marginTop: 5,
                    zIndex: 1,
                    width: 120
                }}
            >
                新增设备模型
            </Button>
            <DeviceModelCreateForm
                visible={deviceModelCreateFormVisible}
                onCreate={onDeviceTypeCreate}
                onCancel={() => {
                    setDeviceModelCreateFormVisible(false);
                }}
            />
            <List
                itemLayout="vertical"
                size="large"
                loading={loading}
                pagination={{
                    onChange: page => {
                        setPagination({
                            current: page,
                        })
                    },
                    pageSize: 7,
                    total: data?.deviceModels.totalCount
                }}
                dataSource={data?.deviceModels.edges}
                renderItem={(item: any) => (
                    <List.Item
                        key={item.id}
                        className={props.currentSelectID === item.id ? "pdc-card-selected" : "pdc-card-default"}
                        onClick={() => props.setCurrentSelectItem(item)}
                    >
                        <div style={{ display: "flex", flexDirection: "column", textAlign: "left" }}>
                            <Tag color="geekblue" style={{ width: "fit-content" }}>{DeviceTypeMap.get(item.deviceType)}</Tag>
                            <div style={{ display: "flex", flexDirection: "row", marginTop: 10 }}>
                                <div style={{ width: 60 }}>ID：{item.id} </div>
                                <div >名称：{item.name}</div>
                            </div>
                        </div>

                    </List.Item>
                )
                }
            />
        </div >)

}