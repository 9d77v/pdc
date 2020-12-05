import { List, Button, Tag, Badge } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { DeviceCreateForm } from './DeviceCreateForm';
import { useMutation, useQuery } from '@apollo/react-hooks';
import "src/styles/card.less"
import { EditOutlined } from '@ant-design/icons';
import { DeviceUpdateForm } from './DeviceUpdateForm';
import { pb } from 'src/pb/compiled';
import useWebSocket from 'react-use-websocket';
import { iotHealthSocketURL } from 'src/utils/ws_client';
import { blobToArrayBuffer } from 'src/utils/file';
import { IDevice, INewDevice, IUpdateDevice } from 'src/models/device';
import { ADD_DEVICE, UPDATE_DEVICE } from 'src/gqls/device/mutation';
import { LIST_DEVICE } from 'src/gqls/device/query';

interface IDeviceListProps {
    currentSelectID: number
    setCurrentSelectItem: (item: IDevice) => void
}

export const DeviceList = (props: IDeviceListProps) => {
    const [dataResource, setDataResource] = useState<any[]>([])
    const [healthMap, setHealthMap] = useState<Map<number, pb.Health>>(new Map<number, pb.Health>())
    const updateHealthCallback: any = useRef();
    const [deviceCreateFormVisible, setDeviceCreateFormVisible] = useState(false)
    const [deviceUpdateFormVisible, setDeviceUpdateFormVisible] = useState(false)
    const [updateDeviceData, setUpdateDeviceData] = useState({
        id: 0,
        name: "",
        ip: "",
        port: 0,
        username: "",
        password: ""
    })
    const [addDevice] = useMutation(ADD_DEVICE);
    const [updateDevice] = useMutation(UPDATE_DEVICE);
    const [pagination, setPagination] = useState({
        current: 1,
    })
    const { loading, data, refetch, } = useQuery(LIST_DEVICE,
        {
            variables: {
                searchParam: {
                    page: pagination.current,
                    pageSize: 7,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                }
            }
        });

    const onDeviceCreate = async (values: INewDevice) => {
        await addDevice({
            variables: {
                "input": {
                    "name": values.name,
                    "deviceModelID": values.deviceModelID,
                    "ip": values.ip,
                    "port": values.port,
                    "username": values.username,
                    "password": values.password
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
                    "name": values.name,
                    "ip": values.ip,
                    "port": values.port,
                    "username": values.username,
                    "password": values.password
                }
            }
        });
        setDeviceUpdateFormVisible(false);
        await refetch()
    };

    useEffect(() => {
        const newDataResource: any[] = []
        const tempData = data?.devices.edges || []
        for (let element of tempData) {
            let t: any = {
                id: element.id,
                deviceModelName: element.deviceModelName,
                deviceModelID: element.deviceModelID,
                ip: element.ip,
                port: element.port,
                name: element.name,
                health: null,
                username: element.username,
                password: element.password,
            }
            newDataResource.push(t)
        }
        setDataResource(newDataResource)
    }, [data])

    const {
        sendMessage,
        lastMessage,
    } = useWebSocket(iotHealthSocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        share: true,
        queryParams: {
            'token': localStorage.getItem('accessToken') || "",
        },
        reconnectAttempts: 720
    });
    useEffect(() => {
        const tempData = data?.devices.edges || []
        if (tempData.length > 0) {
            let subscribeStr = ""
            for (const d of tempData) {
                subscribeStr += d.id + ";"
            }
            sendMessage(subscribeStr);
        }
    }, [data, sendMessage])

    useEffect(() => {
        if (lastMessage) {
            blobToArrayBuffer(lastMessage.data).then((d: any) => {
                const msg = pb.Health.decode(new Uint8Array(d))
                setHealthMap(t => t.set(msg.DeviceID, msg))
            })
        }
    }, [lastMessage])

    const callBack = () => {
        if (healthMap.size > 0) {
            for (let element of dataResource) {
                const msg = healthMap.get(Number(element.id))
                if (msg) {
                    element.health = msg.Value
                    element.updatedAt = msg.ActionTime?.seconds
                }
            }
            setHealthMap(new Map<number, pb.Health>())
        }
    }

    useEffect(() => {
        updateHealthCallback.current = callBack;
        return () => { };
    })

    useEffect(() => {
        const tick = () => {
            updateHealthCallback.current()
        }
        const timer: any = setInterval(tick, 1000)
        return () => {
            clearInterval(timer);
        }
    }, [])

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
                dataSource={dataResource}
                renderItem={(item: IDevice) => {
                    let status: any
                    let statusStr: string
                    switch (item.health) {
                        case 0:
                            status = "error"
                            statusStr = "离线"
                            break
                        case 1:
                            status = "processing"
                            statusStr = "在线"
                            break
                        default:
                            status = "default"
                            statusStr = "未知"
                    }
                    return (
                        <List.Item
                            key={item.id}
                            actions={[
                                <div onClick={
                                    () => {
                                        setUpdateDeviceData({
                                            "id": item.id,
                                            "name": item.name,
                                            "ip": item.ip,
                                            "port": item.port,
                                            "username": item.username,
                                            "password": item.password
                                        })
                                        setDeviceUpdateFormVisible(true)
                                    }
                                }><EditOutlined />编辑</div>
                            ]}
                            className={props.currentSelectID === item.id ? "pdc-card-selected" : "pdc-card-default"}
                        >
                            <div style={{ display: "flex", flexDirection: "column", textAlign: "left" }}
                                onClick={() => props.setCurrentSelectItem(item)}>
                                <span>
                                    <Tag color="geekblue" style={{ width: "fit-content" }}>{item.deviceModelName}</Tag>
                                    <Badge status={status} />
                                    {statusStr}
                                </span>
                                <div style={{ display: "flex", flexDirection: "row", marginTop: 10 }}>
                                    <div style={{ width: 60 }}>ID：{item.id} </div>
                                    <div >名称：{item.name}</div>
                                </div>
                            </div>

                        </List.Item>
                    )
                }}
            />
        </div >)
}
