import { Table, Button, message, } from 'antd'
import { useState, useEffect } from 'react'
import { Img } from 'src/components'
import { useQuery } from '@apollo/react-hooks'
import { BookshelfCreateForm } from './BookshelfCreateForm'
import { useMutation } from '@apollo/react-hooks'
import dayjs from 'dayjs'
import { BookshelfUpdateForm } from './BookshelfUpdateForm'
import { TablePaginationConfig } from 'antd/lib/table'
import Search from 'antd/lib/input/Search'
import { ADD_BOOKSHELF, UPDATE_BOOKSHELF } from 'src/gqls/book/bookshelf.mutation'
import { LIST_BOOKSHELF } from 'src/gqls/book/bookshelf.query'
import { IUpdateBookshelf, IBookshelf } from 'src/module/book/bookshelf.model'

export default function BookshelfTable() {
    const [visible, setVisible] = useState(false)
    const [updateBookshelfVisible, setUpdateBookshelfVisible] = useState(false)
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [updateBookshelfData, setUpdateBookshelfData] = useState<IUpdateBookshelf>({
        id: 0,
        name: "",
        cover: "",
        layerNum: 0,
        partitionNum: 0,
    })
    const [keyword, setKeyword] = useState("")
    const [addBookshelf] = useMutation(ADD_BOOKSHELF)
    const [updateBookshelf] = useMutation(UPDATE_BOOKSHELF)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_BOOKSHELF,
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

    const onBookshelfCreate = async (values: IBookshelf) => {
        await addBookshelf({
            variables: {
                "input": {
                    "name": values.name,
                    "cover": values.cover,
                    "layerNum": values.layerNum,
                    "partitionNum": values.partitionNum,
                }
            }
        })
        setVisible(false)
        await refetch()
    }

    const onBookshelfUpdate = async (values: IUpdateBookshelf) => {
        await updateBookshelf({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name,
                    "cover": values.cover,
                    "layerNum": values.layerNum,
                    "partitionNum": values.partitionNum,
                }
            }
        })
        setUpdateBookshelfVisible(false)
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
                const newEdges = fetchMoreResult ? fetchMoreResult.bookshelfs.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.bookshelfs.totalCount : 0
                setPagination({
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                })
                return newEdges.length
                    ? {
                        bookshelfs: {
                            __typename: previousResult.bookshelfs.__typename,
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
        { title: 'ID', dataIndex: 'id', key: 'id', width: 80, fixed: "left" as const },
        { title: '书架名', dataIndex: 'name', key: 'name', width: 100, fixed: "left" as const },
        {
            title: '图片', dataIndex: 'cover', key: 'cover', width: 100, fixed: "left" as const,
            render: (value: string) => <Img src={value ? value : ''} width={40} height={53.5} />
        }, { title: '层数', dataIndex: 'layerNum', key: 'layerNum', width: 100 },
        { title: '分区数', dataIndex: 'partitionNum', key: 'partitionNum', width: 100 },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 170,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 170,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const, width: 120,
            render: (value: any, record: any) =>
                <span><Button
                    onClick={() => {
                        setUpdateBookshelfData({
                            "id": record.id,
                            "name": record.name,
                            "cover": record.cover,
                            "layerNum": record.layerNum,
                            "partitionNum": record.partitionNum,
                        })
                        setUpdateBookshelfVisible(true)
                    }}>编辑书架</Button>
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
                新增书架
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <BookshelfCreateForm
                visible={visible}
                onCreate={onBookshelfCreate}
                onCancel={() => {
                    setVisible(false)
                }}
            />
            <BookshelfUpdateForm
                visible={updateBookshelfVisible}
                data={updateBookshelfData}
                onUpdate={onBookshelfUpdate}
                onCancel={() => {
                    setUpdateBookshelfVisible(false)
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
                    total: data ? data.bookshelfs.totalCount : 0
                }}
                dataSource={data ? data.bookshelfs.edges : []}
            />
        </div>

    )
}
