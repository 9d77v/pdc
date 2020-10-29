import { Table } from 'antd';
import React from 'react'
import dayjs from 'dayjs';
interface IAttributeTableProps {
    id: number
    data: any[]
}

export default function AttributeTable(props: IAttributeTableProps) {
    const { data } = props

    const columns = [
        { title: 'id', dataIndex: 'id', key: 'id' },
        { title: '键', dataIndex: 'key', key: 'key' },
        { title: '名称', dataIndex: 'name', key: 'name' },
        { title: '值', dataIndex: 'value', key: 'value' },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        }
    ];
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
                    hideOnSinglePage: true,
                    total: data.length
                }}
                dataSource={data}
            />
        </div>
    );
}
