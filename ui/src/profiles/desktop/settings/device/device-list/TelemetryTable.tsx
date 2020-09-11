import { Table } from 'antd';
import React, { useEffect, useState } from 'react'

import moment from 'moment';
import useWebSocket from 'react-use-websocket';
import { deviceTelemetryPrefix, iotSocketURL } from '../../../../../utils/ws_client';
import { pb } from '../../../../../pb/compiled';
import { blobToArrayBuffer } from '../../../../../utils/file';
interface ITelemetryTableProps {
    id: number
    data: any[]
}

export default function TelemetryTable(props: ITelemetryTableProps) {
    const { id, data } = props
    const [dataResource, setDataResource] = useState<any[]>([])
    const columns = [
        { title: 'id', dataIndex: 'id', key: 'id' },
        { title: '键', dataIndex: 'key', key: 'key' },
        { title: '名称', dataIndex: 'name', key: 'name' },
        {
            title: '值', dataIndex: 'value', key: 'value', render: (value: number, record: any) => {
                return value ? (record.factor * value).toFixed(record.scale) : "-"
            }

        },
        { title: '单位', dataIndex: 'unit', key: 'unit' },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => value === null ? "-" : moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        }
    ];

    useEffect(() => {
        const newDataResource: any[] = []
        for (let element of data) {
            let t: any = {
                createdAt: element.createdAt,
                id: element.id,
                key: element.key,
                name: element.name,
                unit: element.unit,
                factor: element.factor,
                scale: element.scale,
                updatedAt: null,
                value: element.value
            }
            newDataResource.push(t)
        }
        setDataResource(newDataResource)
    }, [data])

    const {
        sendMessage,
        lastMessage,
    } = useWebSocket(iotSocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        share: true,
    });
    useEffect(() => {
        sendMessage(deviceTelemetryPrefix + "." + id.toString() + ".*");
    }, [id, sendMessage])

    useEffect(() => {
        if (lastMessage) {
            blobToArrayBuffer(lastMessage.data).then((d: any) => {
                const msg = pb.Telemetry.decode(new Uint8Array(d))
                for (let element of dataResource) {
                    if (Number(element.id) === msg.ID) {
                        element.value = msg.Value
                        element.updatedAt = msg.ActionTime?.seconds
                        break
                    }
                }
            })
        }
    }, [lastMessage, dataResource])
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            width: "100%",
            backgroundColor: "#fff",
            padding: "0px 10px 10px 10px"
        }}>
            <Table
                rowKey={record => record.id}
                columns={columns}
                bordered
                pagination={{
                    pageSize: 5,
                    total: dataResource.length
                }}
                dataSource={dataResource}
            />
        </div>
    )
}
