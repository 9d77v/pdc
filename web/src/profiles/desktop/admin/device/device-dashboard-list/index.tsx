import { List, Button, Popconfirm, Tag, Modal } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { DeviceDashboardCreateForm } from './DeviceDashboardCreateForm';
import { ADD_DEVICE_DASHBOARD, ADD_DEVICE_DASHBOARD_CAMERA, ADD_DEVICE_DASHBOARD_TELEMETRY, DELETE_DEVICE_DASHBOARD, LIST_DEVICE_DASHBOARD, REMOVE_DEVICE_DASHBOARD_CAMERA, REMOVE_DEVICE_DASHBOARD_TELEMETRY, UPDATE_DEVICE_DASHBOARD } from 'src/consts/device.gql';
import { useMutation, useQuery } from '@apollo/react-hooks';
import "src/styles/card.less"
import { DeleteOutlined, EditOutlined, FileAddOutlined, VideoCameraAddOutlined } from '@ant-design/icons';
import { DeviceDashboardUpdateForm } from './DeviceDashboardUpdateForm';
import { pb } from 'src/pb/compiled';
import useWebSocket from 'react-use-websocket';
import { iotTelemetrySocketURL } from 'src/utils/ws_client';
import { blobToArrayBuffer } from 'src/utils/file';
import { DeviceDashboardTelemetryAddForm } from './DeviceDashboardTelemetryAddForm';
import { DeviceDashboardCameraAddForm } from './DeviceDashboardCameraAddForm';
import { IDeviceDashboard, IDeviceDashboardCamera, IDeviceDashboardTelemetry, INewDeviceDashboard, INewDeviceDashboardCamera, INewDeviceDashboardTelemetry, IUpdateDeviceDashboard } from 'src/models/device';
import CameraCard from 'src/profiles/common/device/CameraPicture';
import { LivePlayer } from 'src/components/videoplayer/LivePlayer';

export default function DeviceDashboardList() {
    const [dataResource, setDataResource] = useState<any[]>([])
    const [telemetryMap, setTelemetryMap] = useState<Map<number, pb.Telemetry>>(new Map<number, pb.Telemetry>())
    const updateTelemetryCallback: any = useRef();
    const [deviceDashboardCreateFormVisible, setDeviceDashboardCreateFormVisible] = useState(false)
    const [deviceDashboardUpdateFormVisible, setDeviceDashboardUpdateFormVisible] = useState(false)
    const [deviceDashboardTelemetryAddFormVisible, setDeviceDashboardTelemetryAddFormVisible] = useState(false)
    const [deviceDashboardCameraAddFormVisible, setDeviceDashboardCameraAddFormVisible] = useState(false)

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
    const [addDeviceDashboardCamera] = useMutation(ADD_DEVICE_DASHBOARD_CAMERA)
    const [removeDeviceDashboardCamera] = useMutation(REMOVE_DEVICE_DASHBOARD_CAMERA)

    const [currentDeviceDashboard, setCurrentDeviceDashboard] = useState<{
        id: number,
        deviceType: number,
        telemetries: IDeviceDashboardTelemetry[],
        cameras: IDeviceDashboardCamera[]
    }>({
        id: 0,
        deviceType: 0,
        telemetries: [],
        cameras: []
    })

    const [currentCamera, setCurrentCamera] = useState({
        url: "",
        title: "",
        visible: false
    })

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
                    "deviceType": values.deviceType
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
                    "deviceDashboardID": currentDeviceDashboard.id,
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

    const onDeviceDashboardCameraAdd = async (values: INewDeviceDashboardCamera) => {
        await addDeviceDashboardCamera({
            variables: {
                "input": {
                    "deviceDashboardID": currentDeviceDashboard.id,
                    "deviceIDs": values.deviceIDs,
                }
            }
        });
        setDeviceDashboardCameraAddFormVisible(false)
        await refetch()
    }

    const onDeviceDashboardCameraRemove = async (id: number) => {
        await removeDeviceDashboardCamera({
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
                for (const t of element.telemetries || []) {
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
                let cameras: any[] = []
                for (const t of element.cameras || []) {
                    cameras.push({
                        id: t.id,
                        deviceID: t.deviceID,
                        deviceName: t.deviceName
                    })
                }
                let t: any = {
                    createdAt: element.createdAt,
                    id: element.id,
                    name: element.name,
                    isVisible: element.isVisible,
                    deviceType: element.deviceType,
                    telemetries: telemetries,
                    cameras: cameras
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

    const showActionButton = (item: {
        id: number,
        deviceType: number,
        telemetries: IDeviceDashboardTelemetry[],
        cameras: IDeviceDashboardCamera[]
    }) => {
        const actionButtonMap = new Map<number, JSX.Element>([
            [0, <div onClick={
                () => {
                    setDeviceDashboardTelemetryAddFormVisible(true)
                    setCurrentDeviceDashboard({
                        id: item.id,
                        deviceType: item.deviceType,
                        telemetries: item.telemetries,
                        cameras: item.cameras
                    })
                }
            }><FileAddOutlined />添加遥测</div>],
            [1, <div onClick={
                () => {
                    setDeviceDashboardCameraAddFormVisible(true)
                    setCurrentDeviceDashboard({
                        id: item.id,
                        deviceType: item.deviceType,
                        telemetries: item.telemetries,
                        cameras: item.cameras
                    })
                }
            }><VideoCameraAddOutlined />添加摄像头</div>]
        ])
        return actionButtonMap.get(item.deviceType)
    }

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
                existData={currentDeviceDashboard.telemetries}
                onCreate={onDeviceDashboardTelemetryAdd}
                onCancel={() => {
                    setDeviceDashboardTelemetryAddFormVisible(false)
                }}
            />
            <DeviceDashboardCameraAddForm
                visible={deviceDashboardCameraAddFormVisible}
                existData={currentDeviceDashboard.cameras}
                onCreate={onDeviceDashboardCameraAdd}
                onCancel={() => {
                    setDeviceDashboardCameraAddFormVisible(false)
                }}
            />
            <Modal
                visible={currentCamera.visible}
                title={currentCamera.title}
                footer={null}
                destroyOnClose={true}
                getContainer={false}
                width={1020}
                onCancel={
                    () => {
                        setCurrentCamera({
                            title: "",
                            url: "",
                            visible: false,
                        })
                    }
                }
            >  <LivePlayer
                    url={currentCamera.url}
                    height={540}
                    width={960}
                /></Modal>
            <List
                itemLayout="vertical"
                size="large"
                loading={loading}
                dataSource={dataResource}
                renderItem={(item: IDeviceDashboard) => {
                    const telemetryList = <List
                        style={{ minHeight: 100 }}
                        itemLayout="horizontal"
                        dataSource={item.deviceType === 0 ? item.telemetries : item.cameras}
                        renderItem={(t: any) => {
                            switch (item.deviceType) {
                                case 0:
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
                                case 1:
                                    return (
                                        <List.Item
                                            actions={[<div onClick={() => onDeviceDashboardCameraRemove(t.id)}
                                            >移除</div>]}
                                        >
                                            <div style={{ width: "100%" }} onClick={() => {
                                                setCurrentCamera({
                                                    title: t.deviceName,
                                                    url: `/hls/stream${t.deviceID}.m3u8`,
                                                    visible: true,
                                                })
                                            }} >
                                                <div style={{ textAlign: "left", margin: 10 }}>{t.deviceName}</div>
                                                <CameraCard
                                                    border={"1px solid grey"}
                                                    minHeight={100}
                                                    deviceID={t.deviceID} />
                                            </div>
                                        </List.Item>)
                                default:
                                    return (<div />)
                            }

                        }
                        }
                    />
                    return (
                        <List.Item
                            key={item.id}
                            actions={[
                                showActionButton(item),
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
