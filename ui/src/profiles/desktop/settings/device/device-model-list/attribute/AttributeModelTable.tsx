import { Table, Button } from 'antd';
import React, { useState } from 'react'

import { useMutation } from '@apollo/react-hooks';
import moment from 'moment';
import { AttributeModelCreateForm, INewAttributeModel } from './AttributeModelCreateForm';
import { ADD_ATTRIBUTE_MODEL, UPDATE_ATTRIBUTE_MODEL } from '../../../../../../consts/device.gql';
import { IUpdateAttributeModel, AttributeModelUpdateForm } from './AttributeModelUpdateForm';


interface IAttributeModelTableProps {
    id: number
    data: any[]
    refetch: () => void
}

export default function AttributeModelTable(props: IAttributeModelTableProps) {
    const { id, data, refetch } = props
    const [attributeModelCreateFormVisible, setAttributeModelCreateFormVisible] = useState(false);
    const [attributeModelUpdateFormVisible, setAttributeModelUpdateFormVisible] = useState(false)
    const [updateAttributeModelData, setUpdateAttributeModelData] = useState({
        id: 0,
        key: "",
        name: ""
    })
    const [addAttributeModel] = useMutation(ADD_ATTRIBUTE_MODEL);
    const [updateAttributeModel] = useMutation(UPDATE_ATTRIBUTE_MODEL)

    const onAttributeModelCreate = async (values: INewAttributeModel) => {
        await addAttributeModel({
            variables: {
                "input": {
                    "deviceModelID": id,
                    "key": values.key,
                    "name": values.name
                }
            }
        });
        setAttributeModelCreateFormVisible(false);
        refetch()
    };

    const onAttributeModelUpdate = async (values: IUpdateAttributeModel) => {
        await updateAttributeModel({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name
                }
            }
        });
        setAttributeModelUpdateFormVisible(false);
        await refetch()
    };

    const columns = [
        { title: '键', dataIndex: 'key', key: 'key' },
        { title: '名称', dataIndex: 'name', key: 'name' },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const,
            render: (value: any, record: any) =>
                <span><Button
                    onClick={() => {
                        setUpdateAttributeModelData({
                            "id": record.id,
                            "key": record.key,
                            "name": record.name
                        })
                        setAttributeModelUpdateFormVisible(true)
                    }}>编辑属性</Button>
                </span >
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
                    setAttributeModelCreateFormVisible(true);
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 100 }}
            >
                新增属性
            </Button>
            <AttributeModelCreateForm
                visible={attributeModelCreateFormVisible}
                onCreate={onAttributeModelCreate}
                onCancel={() => {
                    setAttributeModelCreateFormVisible(false);
                }}
            />
            <AttributeModelUpdateForm
                visible={attributeModelUpdateFormVisible}
                data={updateAttributeModelData}
                onUpdate={onAttributeModelUpdate}
                onCancel={() => {
                    setAttributeModelUpdateFormVisible(false);
                }}
            />
            <Table
                rowKey={record => record.id}
                columns={columns}
                bordered
                pagination={{
                    pageSize: 5,
                    hideOnSinglePage: true,
                    total: data.length
                }}
                dataSource={data}
            />
        </div>
    );
}
