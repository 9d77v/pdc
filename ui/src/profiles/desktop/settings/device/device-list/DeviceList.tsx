import { List, Button, Tag } from 'antd';
import React, { useState } from 'react';
import { DeviceCreateForm, INewDevice } from './DeviceCreateForm';
import { ADD_DEVICE, LIST_DEVICE, UPDATE_DEVICE } from '../../../../../consts/device.gql';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { IDevice } from '../../../../../consts/consts';
import "../../../../../style/card.less"
import { DeleteOutlined } from '@ant-design/icons';
import { IUpdateDevice, DeviceUpdateForm } from './DeviceUpdateForm';

interface IDeviceListProps {
    currentSelectID: number
    setCurrentSelectItem: (item: IDevice) => void
}

export const DeviceList = (props: IDeviceListProps) => {
    const [deviceCreateFormVisible, setDeviceCreateFormVisible] = useState(false)
    const [deviceUpdateFormVisible, setDeviceUpdateFormVisible] = useState(false)
    const [updateDeviceData, setUpdateDeviceData] = useState({
        id: 0,
        name: ""
    })
    const [addDevice] = useMutation(ADD_DEVICE);
    const [updateDevice] = useMutation(UPDATE_DEVICE);
    const [pagination, setPagination] = useState({
        current: 1,
    })
    const { loading, data, refetch, } = useQuery(LIST_DEVICE,
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

    const onDeviceCreate = async (values: INewDevice) => {
        await addDevice({
            variables: {
                "input": {
                    "name": values.name,
                    "deviceModelID": values.deviceModelID
                }
            }
        });
        setDeviceCreateFormVisible(false);
        await refetch()
    };

    const onDeviceUpdate = async (values: IUpdateDevice) => {
        await updateDevice({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name
                }
            }
        });
        setDeviceUpdateFormVisible(false);
        await refetch()
    };
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setDeviceCreateFormVisible(true)
                }}
                style={{
                    float: 'left',
                    marginLeft: 20,
                    marginTop: 5,
                    zIndex: 1,
                    width: 120
                }}
            >
                新增设备
            </Button>
            <DeviceCreateForm
                visible={deviceCreateFormVisible}
                onCreate={onDeviceCreate}
                onCancel={() => {
                    setDeviceCreateFormVisible(false);
                }}
            />
            <DeviceUpdateForm
                data={updateDeviceData}
                visible={deviceUpdateFormVisible}
                onUpdate={onDeviceUpdate}
                onCancel={() => {
                    setDeviceUpdateFormVisible(false);
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
                    total: data?.devices.totalCount
                }}
                dataSource={data?.devices.edges}
                renderItem={(item: IDevice) => (
                    <List.Item
                        key={item.id}
                        actions={[
                            <div onClick={
                                () => {
                                    setUpdateDeviceData({
                                        "id": item.id,
                                        "name": item.name
                                    })
                                    setDeviceUpdateFormVisible(true)
                                }
                            }><DeleteOutlined />编辑</div>
                        ]}
                        className={props.currentSelectID === item.id ? "pdc-card-selected" : "pdc-card-default"}
                    >
                        <div style={{ display: "flex", flexDirection: "column", textAlign: "left" }}
                            onClick={() => props.setCurrentSelectItem(item)}>
                            <Tag color="geekblue" style={{ width: "fit-content" }}>{item.deviceModelName}</Tag>
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
