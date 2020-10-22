import { List, Button, Popconfirm, Tag } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { DeviceDashboardCreateForm, INewDeviceDashboard } from './DeviceDashboardCreateForm';
import { ADD_DEVICE_DASHBOARD, ADD_DEVICE_DASHBOARD_TELEMETRY, DELETE_DEVICE_DASHBOARD, LIST_DEVICE_DASHBOARD, REMOVE_DEVICE_DASHBOARD_TELEMETRY, UPDATE_DEVICE_DASHBOARD } from '../../../../../consts/device.gql';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { IDeviceDashboard } from '../../../../../consts/consts';
import "../../../../../style/card.less"
import { DeleteOutlined, EditOutlined, FileAddOutlined } from '@ant-design/icons';
import { IUpdateDeviceDashboard, DeviceDashboardUpdateForm } from './DeviceDashboardUpdateForm';
import { pb } from '../../../../../pb/compiled';
import useWebSocket from 'react-use-websocket';
import { iotTelemetrySocketURL } from '../../../../../utils/ws_client';
import { blobToArrayBuffer } from '../../../../../utils/file';
import { DeviceDashboardTelemetryAddForm, INewDeviceDashboardTelemetry } from './DeviceDashboardTelemetryAddForm';

export default function DeviceDashboardList() {
    const [dataResource, setDataResource] = useState<any[]>([])
    const [telemetryMap, setTelemetryMap] = useState<Map<number, pb.Telemetry>>(new Map<number, pb.Telemetry>())
    const updateTelemetryCallback: any = useRef();
    const [deviceDashboardCreateFormVisible, setDeviceDashboardCreateFormVisible] = useState(false)
    const [deviceDashboardUpdateFormVisible, setDeviceDashboardUpdateFormVisible] = useState(false)
    const [deviceDashboardTelemetryAddFormVisible, setDeviceDashboardTelemetryAddFormVisible] = useState(false)
    const [updateDeviceDashboardData, setUpdateDeviceDashboardData] = useState({
        id: 0,
        name: "",
        isVisible: false
    })
    const [addDeviceDashboard] = useMutation(ADD_DEVICE_DASHBOARD)
    const [updateDeviceDashboard] = useMutation(UPDATE_DEVICE_DASHBOARD)
    const [deleteDeviceDashboard] = useMutation(DELETE_DEVICE_DASHBOARD)
    const [addDeviceDashboardTelemetry] = useMutation(ADD_DEVICE_DASHBOARD_TELEMETRY)
    const [removeDeviceDashboardTelemetry] = useMutation(REMOVE_DEVICE_DASHBOARD_TELEMETRY)

    const [currentDeviceDashboardID, setCurrentDeviceDashboardID] = useState(0)
    const { loading, data, refetch, } = useQuery(LIST_DEVICE_DASHBOARD,
        {
            variables: {
                sorts: [{
                    field: 'id',
                    isAsc: true
                }]
            },
            fetchPolicy: "cache-and-network"
        })

    const onDeviceDashboardCreate = async (values: INewDeviceDashboard) => {
        await addDeviceDashboard({
            variables: {
                "input": {
                    "name": values.name,
                    "isVisible": values.isVisible,
                }
            }
        });
        setDeviceDashboardCreateFormVisible(false)
        await refetch()
    };

    const onDeviceDashboardUpdate = async (values: IUpdateDeviceDashboard) => {
        await updateDeviceDashboard({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name,
                    "isVisible": values.isVisible,
                }
            }
        })
        setDeviceDashboardUpdateFormVisible(false)
        await refetch()
    }

    const onDeviceDashboardTelemetryAdd = async (values: INewDeviceDashboardTelemetry) => {
        await addDeviceDashboardTelemetry({
            variables: {
                "input": {
                    "deviceDashboardID": currentDeviceDashboardID,
                    "telemetryIDs": values.telemetryIDs,
                }
            }
        });
        setDeviceDashboardTelemetryAddFormVisible(false)
        await refetch()
    }
    const onDeviceDashboardTelemetryRemove = async (id: number) => {
        await removeDeviceDashboardTelemetry({
            variables: {
                "ids": [id]
            }
        });
        await refetch()
    }

    useEffect(() => {
        if (data) {
            const newDataResource: any[] = []
            for (const element of data.deviceDashboards.edges) {
                let telemetries: any[] = []
                for (const t of element.telemetries) {
                    telemetries.push({
                        createdAt: t.createdAt,
                        id: t.id,
                        telemetryID: t.telemetryID,
                        deviceID: t.deviceID,
                        key: t.key,
                        name: t.name,
                        unit: t.unit,
                        unitName: t.unitName,
                        factor: t.factor,
                        scale: t.scale,
                        updatedAt: null,
                        value: t.value,
                        deviceName: t.deviceName
                    })
                }
                let t: any = {
                    createdAt: element.createdAt,
                    id: element.id,
                    name: element.name,
                    isVisible: element.isVisible,
                    telemetries: telemetries
                }
                newDataResource.push(t)
            }
            setDataResource(newDataResource)
        }
    }, [data])

    const {
        sendMessage,
        lastMessage,
    } = useWebSocket(iotTelemetrySocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        share: true,
        queryParams: {
            'token': localStorage.getItem('accessToken') || "",
        },
        reconnectAttempts: 720
    })

    useEffect(() => {
        let telemetries: string[] = []
        for (let element of data ? data.deviceDashboards.edges : []) {
            for (let t of element.telemetries) {
                telemetries.push(t.deviceID + "." + t.telemetryID)
            }
        }
        if (telemetries.length > 0) {
            sendMessage(telemetries.join(";"))
        }
    }, [sendMessage, data])

    useEffect(() => {
        if (lastMessage) {
            blobToArrayBuffer(lastMessage.data).then((d: any) => {
                const msg = pb.Telemetry.decode(new Uint8Array(d))
                setTelemetryMap(t => t.set(msg.ID, msg))
            })
        }
    }, [lastMessage])

    const callBack = () => {
        if (telemetryMap.size > 0) {
            for (let element of dataResource) {
                for (let t of element.telemetries) {
                    const msg = telemetryMap.get(Number(t.telemetryID))
                    if (msg) {
                        t.value = msg.Value
                        t.updatedAt = msg.ActionTime?.seconds
                    }
                }
            }
            setTelemetryMap(new Map<number, pb.Telemetry>())
        }
    }

    useEffect(() => {
        updateTelemetryCallback.current = callBack;
        return () => { };
    })

    useEffect(() => {
        const tick = () => {
            updateTelemetryCallback.current()
        }
        const timer: any = setInterval(tick, 1000)
        return () => {
            clearInterval(timer)
        }
    }, [])
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setDeviceDashboardCreateFormVisible(true)
                }}
                style={{
                    float: 'left',
                    marginLeft: 20,
                    marginTop: 5,
                    zIndex: 1,
                    width: 140
                }}
            >
                新增设备仪表盘
            </Button>
            <DeviceDashboardCreateForm
                visible={deviceDashboardCreateFormVisible}
                onCreate={onDeviceDashboardCreate}
                onCancel={() => {
                    setDeviceDashboardCreateFormVisible(false)
                }}
            />
            <DeviceDashboardUpdateForm
                data={updateDeviceDashboardData}
                visible={deviceDashboardUpdateFormVisible}
                onUpdate={onDeviceDashboardUpdate}
                onCancel={() => {
                    setDeviceDashboardUpdateFormVisible(false)
                }}
            />
            <DeviceDashboardTelemetryAddForm
                visible={deviceDashboardTelemetryAddFormVisible}
                onCreate={onDeviceDashboardTelemetryAdd}
                onCancel={() => {
                    setDeviceDashboardTelemetryAddFormVisible(false)
                }}
            />
            <List
                itemLayout="vertical"
                size="large"
                loading={loading}
                dataSource={dataResource}
                renderItem={(item: IDeviceDashboard) => {
                    const telemetryList = <List
                        style={{ minHeight: 100 }}
                        itemLayout="horizontal"
                        dataSource={item.telemetries}
                        renderItem={t => {
                            const value = t.value === null ? "-" : (t.factor * (t.value || 0)).toFixed(t.scale)
                            return (
                                <List.Item
                                    actions={[<div onClick={() => onDeviceDashboardTelemetryRemove(t.id)}
                                    >移除</div>]}
                                >
                                    <Tag color="geekblue" style={{ width: 100 }}>{t.deviceName}</Tag>
                                    <div key={t.id}>
                                        {t.name}: {value}{t.unit}</div>
                                </List.Item>
                            )
                        }
                        }
                    />
                    return (
                        <List.Item
                            key={item.id}
                            actions={[
                                <div onClick={
                                    () => {
                                        setDeviceDashboardTelemetryAddFormVisible(true)
                                        setCurrentDeviceDashboardID(item.id)
                                    }
                                }><FileAddOutlined />添加遥测</div>,
                                <div onClick={
                                    () => {
                                        setUpdateDeviceDashboardData({
                                            "id": item.id,
                                            "name": item.name,
                                            "isVisible": item.isVisible,
                                        })
                                        setDeviceDashboardUpdateFormVisible(true)
                                    }
                                }><EditOutlined />编辑</div>,
                                <Popconfirm
                                    title="确定要删除该仪表盘吗?"
                                    onConfirm={async () => {
                                        await deleteDeviceDashboard({
                                            variables: {
                                                "id": item.id
                                            }
                                        })
                                        refetch()
                                    }}
                                    onCancel={() => { }}
                                    okText="是"
                                    cancelText="否"
                                >
                                    <div ><DeleteOutlined />删除</div>
                                </Popconfirm>
                            ]}
                            className={"pdc-card-default"}
                            style={{ margin: 20, float: "left", borderRadius: 40 }}
                        >
                            <div style={{
                                height: 400,
                                width: 400,
                                display: "flex",
                                flexDirection: "column",
                                overflowY: "auto"
                            }} >
                                <div >{item.name}</div>
                                {telemetryList}
                            </div>
                        </List.Item>
                    )
                }}
            />
        </div >)
}
