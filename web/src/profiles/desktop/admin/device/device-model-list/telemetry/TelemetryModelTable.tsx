import { Table, Button, Popconfirm } from 'antd';
import React, { useState } from 'react'

import { useMutation } from '@apollo/react-hooks';
import dayjs from 'dayjs';
import { TelemetryModelCreateForm } from './TelemetryModelCreateForm';
import { UPDATE_TELEMETRY_MODEL, ADD_TELEMETRY_MODEL, DELETE_TELEMETRY_MODEL } from 'src/consts/device.gql';
import { TelemetryModelUpdateForm } from './TelemetryModelUpdateForm';
import { INewTelemetryModel, IUpdateTelemetryModel } from 'src/models/device';


interface ITelemetryModelTableProps {
    id: number
    data: any[]
    refetch: () => void
}

export default function TelemetryModelTable(props: ITelemetryModelTableProps) {
    const { id, data, refetch } = props
    const [attributeModelCreateFormVisible, setTelemetryModelCreateFormVisible] = useState(false);
    const [attributeModelUpdateFormVisible, setTelemetryModelUpdateFormVisible] = useState(false);
    const [updateTelemetryModelData, setUpdateTelemetryModelData] = useState({
        id: 0,
        key: "",
        name: "",
        factor: 1,
        unit: "",
        unitName: "",
        scale: 2
    })
    const [addTelemetryModel] = useMutation(ADD_TELEMETRY_MODEL);
    const [updateTelemetryModel] = useMutation(UPDATE_TELEMETRY_MODEL)
    const [deleteTelemetryModel] = useMutation(DELETE_TELEMETRY_MODEL)

    const onTelemetryModelCreate = async (values: INewTelemetryModel) => {
        await addTelemetryModel({
            variables: {
                "input": {
                    "deviceModelID": id,
                    "key": values.key,
                    "name": values.name,
                    "factor": values.factor,
                    "unit": values.unit,
                    "unitName": values.unitName,
                    "scale": values.scale
                }
            }
        });
        setTelemetryModelCreateFormVisible(false);
        refetch()
    };

    const onTelemetryModelUpdate = async (values: IUpdateTelemetryModel) => {
        await updateTelemetryModel({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name,
                    "factor": values.factor,
                    "unit": values.unit,
                    "unitName": values.unitName,
                    "scale": values.scale
                }
            }
        });
        setTelemetryModelUpdateFormVisible(false);
        await refetch()
    };

    const columns = [
        { title: '键', dataIndex: 'key', key: 'key' },
        { title: '名称', dataIndex: 'name', key: 'name' },
        { title: '系数', dataIndex: 'factor', key: 'factor' },
        { title: '单位', dataIndex: 'unit', key: 'unit' },
        { title: '单位名称', dataIndex: 'unitName', key: 'unitName' },
        { title: '小数位数', dataIndex: 'scale', key: 'scale' },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const,
            render: (value: any, record: any) =>
                <div><Button
                    style={{ marginBottom: 8 }}
                    onClick={() => {
                        setUpdateTelemetryModelData({
                            "id": record.id,
                            "key": record.key,
                            "name": record.name,
                            "factor": record.factor,
                            "unit": record.unit,
                            "unitName": record.unitName,
                            "scale": record.scale
                        })
                        setTelemetryModelUpdateFormVisible(true)
                    }}>编辑遥测</Button>
                    <Popconfirm
                        title="确定要删除该遥测吗?"
                        onConfirm={async () => {
                            await deleteTelemetryModel({
                                variables: {
                                    "id": record.id
                                }
                            })
                            refetch()
                        }}
                        onCancel={() => { }}
                        okText="是"
                        cancelText="否"
                    >
                        <Button
                            danger
                        >删除遥测</Button>
                    </Popconfirm>
                </div>
        },
    ];
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            width: "100%",
            backgroundColor: "#fff",
            padding: "0px 10px 10px 10px"
        }}>
            <Button
                type="primary"
                onClick={() => {
                    setTelemetryModelCreateFormVisible(true);
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 100 }}
            >
                新增遥测
            </Button>
            <TelemetryModelCreateForm
                visible={attributeModelCreateFormVisible}
                onCreate={onTelemetryModelCreate}
                onCancel={() => {
                    setTelemetryModelCreateFormVisible(false);
                }}
            />
            <TelemetryModelUpdateForm
                visible={attributeModelUpdateFormVisible}
                data={updateTelemetryModelData}
                onUpdate={onTelemetryModelUpdate}
                onCancel={() => {
                    setTelemetryModelUpdateFormVisible(false);
                }}
            />
            <Table
                rowKey={record => record.id}
                columns={columns}
                bordered
                pagination={{
                    pageSize: 10,
                    total: data.length
                }}
                dataSource={data}
            />
        </div>
    );
}
