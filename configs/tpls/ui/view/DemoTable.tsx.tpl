import { Table, Button, message } from 'antd'
import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { [[.Name]]CreateForm } from './[[.Name]]CreateForm'
import { useMutation } from '@apollo/react-hooks'
import dayjs from 'dayjs'
import { [[.Name]]UpdateForm } from './[[.Name]]UpdateForm'
import { TablePaginationConfig } from 'antd/lib/table'
import Search from 'antd/lib/input/Search'
import { ADD_[[.TitleName]], UPDATE_[[.TitleName]] } from 'src/gqls/[[.LowerName]]/mutation'
import { LIST_[[.TitleName]] } from 'src/gqls/[[.LowerName]]/query'
import { IUpdate[[.Name]], I[[.Name]] } from 'src/module/[[.LowerName]]/[[.LowerName]].model'

export default function [[.Name]]Table() {
    const [visible, setVisible] = useState(false)
    const [update[[.Name]]Visible, setUpdate[[.Name]]Visible] = useState(false)
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [update[[.Name]]Data, setUpdate[[.Name]]Data] = useState({
        id: 0,
[[range .Columns]][[if eq .Type "Time"]]
        [[.Name]]: dayjs(0),[[else if eq .Type "int"]]
        [[.Name]]: 0,
    [[else]]    
        [[.Name]]: "",
[[end]][[end]]    })
    const [keyword, setKeyword] = useState("")
    const [add[[.Name]]] = useMutation(ADD_[[.TitleName]])
    const [update[[.Name]]] = useMutation(UPDATE_[[.TitleName]])
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_[[.TitleName]],
        {
            variables: {
                searchParam: {
                    page: pagination.current,
                    pageSize: pagination.pageSize,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                },
            },
            fetchPolicy: "cache-and-network"
        })


    const on[[.Name]]Create = async (values: I[[.Name]]) => {
        await add[[.Name]]({
            variables: {
                "input": {
                [[range .Columns]][[if eq .Type "Time"]]
                    "[[.Name]]": values.[[.Name]].unix(),
                [[else]]    
                    "[[.Name]]": values.[[.Name]],
            [[end]][[end]]    }
            }
        })
        setVisible(false)
        await refetch()
    }

    const on[[.Name]]Update = async (values: IUpdate[[.Name]]) => {
        await update[[.Name]]({
            variables: {
                "input": {
                    "id": values.id,[[range .Columns]][[if eq .Type "Time"]]
                    "[[.Name]]": values.[[.Name]].unix(),[[else]]    
                    "[[.Name]]": values.[[.Name]],
                [[end]][[end]]    }
            }
        })
        setUpdate[[.Name]]Visible(false)
        await refetch()
    }

    const onChange = async (pageConfig: TablePaginationConfig) => {
        fetchMore({
            variables: {
                searchParam: {
                    page: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                }
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.[[.LowerName]]s.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.[[.LowerName]]s.totalCount : 0
                setPagination({
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                })
                return newEdges.length
                    ? {
                        [[.LowerName]]s: {
                            __typename: previousResult.[[.LowerName]]s.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult
            }
        })
    }
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])
    const columns = [
        { title: '名称', dataIndex: 'name', key: 'name', width: 200 },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 160,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 160,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const, width: 120,
            render: (value: any, record: IUpdate[[.Name]]) =>
                <span><Button
                    onClick={() => {
                        setUpdate[[.Name]]Data({
                            "id": record.id,
                        [[range .Columns]]    "[[.Name]]": record.[[.Name]],
                    [[end]]    })
                        setUpdate[[.Name]]Visible(true)
                    }}>编辑[[.ModuleName]]</Button>
                </span>
        },
    ]
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setVisible(true)
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 100 }}
            >
                新增[[.ModuleName]]
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <[[.Name]]CreateForm
                visible={visible}
                onCreate={on[[.Name]]Create}
                onCancel={() => {
                    setVisible(false)
                }}
            />
            <[[.Name]]UpdateForm
                visible={update[[.Name]]Visible}
                data={update[[.Name]]Data}
                onUpdate={on[[.Name]]Update}
                onCancel={() => {
                    setUpdate[[.Name]]Visible(false)
                }}
            />
            <Table
                loading={loading}
                rowKey={record => record.id}
                columns={columns}
                scroll={{ x: 1300 }}
                bordered
                onChange={onChange}
                pagination={{
                    ...pagination,
                    total: data ? data.[[.LowerName]]s.totalCount : 0
                }}
                dataSource={data ? data.[[.LowerName]]s.edges : []}
            />
        </div>

    )
}
