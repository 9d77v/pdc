import { List, Button, Tag } from 'antd';
import React, { useState } from 'react';
import { DeviceModelCreateForm, INewDeviceModel } from './DeviceModelCreateForm';
import { ADD_DEVICE_MODEL, LIST_DEVICE_MODEL, UPDATE_DEVICE_MODEL } from '../../../../../consts/device.gql';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { DeviceTypeMap, IDeviceModel } from '../../../../../consts/consts';
import "../../../../../style/card.less"
import { EditOutlined } from '@ant-design/icons';
import { IUpdateDeviceModel, DeviceModelUpdateForm } from './DeviceModelUpdateForm';

interface IDeviceModelListProps {
    currentSelectID: number
    setCurrentSelectItem: (item: IDeviceModel) => void
}

export const DeviceModelList = (props: IDeviceModelListProps) => {
    const [deviceModelCreateFormVisible, setDeviceModelCreateFormVisible] = useState(false)
    const [deviceModelUpdateFormVisible, setDeviceModelUpdateFormVisible] = useState(false)
    const [updateDeviceModelData, setUpdateDeviceModelData] = useState({
        id: 0,
        name: "",
        desc: ""
    })
    const [addDeviceModel] = useMutation(ADD_DEVICE_MODEL);
    const [updateDeviceModel] = useMutation(UPDATE_DEVICE_MODEL);
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

    const onDeviceModelCreate = async (values: INewDeviceModel) => {
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

    const onDeviceModelUpdate = async (values: IUpdateDeviceModel) => {
        await updateDeviceModel({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name,
                    "desc": values.desc,
                }
            }
        });
        setDeviceModelUpdateFormVisible(false);
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
                onCreate={onDeviceModelCreate}
                onCancel={() => {
                    setDeviceModelCreateFormVisible(false);
                }}
            />
            <DeviceModelUpdateForm
                data={updateDeviceModelData}
                visible={deviceModelUpdateFormVisible}
                onUpdate={onDeviceModelUpdate}
                onCancel={() => {
                    setDeviceModelUpdateFormVisible(false);
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
                        actions={[
                            <div onClick={
                                () => {
                                    setUpdateDeviceModelData({
                                        "id": item.id,
                                        "name": item.name,
                                        "desc": item.desc
                                    })
                                    setDeviceModelUpdateFormVisible(true)
                                }
                            }><EditOutlined />编辑</div>
                        ]}
                        className={props.currentSelectID === item.id ? "pdc-card-selected" : "pdc-card-default"}
                    >
                        <div style={{ display: "flex", flexDirection: "column", textAlign: "left" }}
                            onClick={() => props.setCurrentSelectItem(item)}>
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