import { Table, Button, message, Tag } from 'antd'
import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { BookCreateForm } from './BookCreateForm'
import { useMutation } from '@apollo/react-hooks'
import dayjs from 'dayjs'
import { Img } from 'src/components'
import { BookUpdateForm } from './BookUpdateForm'
import { TablePaginationConfig } from 'antd/lib/table'
import Search from 'antd/lib/input/Search'
import { ADD_BOOK, UPDATE_BOOK } from 'src/gqls/book/book.mutation'
import { LIST_BOOK } from 'src/gqls/book/book.query'
import { IUpdateBook, IBook } from 'src/module/book/book.model'
import { useHistory } from 'react-router'
import { AdminPath } from 'src/consts/path'

export default function BookTable() {
    const history = useHistory()
    const [visible, setVisible] = useState(false)
    const [updateBookVisible, setUpdateBookVisible] = useState(false)
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [updateBookID, setUpdateBookID] = useState(0)
    const [keyword, setKeyword] = useState("")
    const [addBook] = useMutation(ADD_BOOK)
    const [updateBook] = useMutation(UPDATE_BOOK)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_BOOK,
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

    const onBookCreate = async (values: IBook) => {
        await addBook({
            variables: {
                "input": {
                    "isbn": values.isbn,
                    "name": values.name,
                    "desc": values.desc,
                    "cover": values.cover,
                    "author": values.author,
                    "translator": values.translator,
                    "publishingHouse": values.publishingHouse,
                    "edition": values.edition,
                    "printedTimes": values.printedTimes,
                    "printedSheets": values.printedSheets,
                    "format": values.format,
                    "wordCount": values.wordCount?.toString(),
                    "pricing": values.pricing?.toString(),
                    "packing": values.packing,
                    "pageSize": values.pageSize,
                    "purchasePrice": values.purchasePrice?.toString(),
                    "purchaseTime": values.purchaseTime ? values.purchaseTime.unix() : 0,
                    "purchaseSource": values.purchaseSource,
                    "bookBorrowUID": values.bookBorrowUID,
                }
            }
        })
        setVisible(false)
        await refetch()
    }

    const onBookUpdate = async (values: IUpdateBook) => {
        await updateBook({
            variables: {
                "input": {
                    "id": values.id,
                    "isbn": values.isbn,
                    "name": values.name,
                    "desc": values.desc,
                    "cover": values.cover,
                    "author": values.author,
                    "translator": values.translator,
                    "publishingHouse": values.publishingHouse,
                    "edition": values.edition,
                    "printedTimes": values.printedTimes,
                    "printedSheets": values.printedSheets,
                    "format": values.format,
                    "wordCount": values.wordCount?.toString(),
                    "pricing": values.pricing?.toString(),
                    "packing": values.packing,
                    "pageSize": values.pageSize,
                    "purchasePrice": values.purchasePrice?.toString(),
                    "purchaseTime": values.purchaseTime ? values.purchaseTime.unix() : 0,
                    "purchaseSource": values.purchaseSource,
                    "bookBorrowUID": values.bookBorrowUID,
                }
            }
        })
        setUpdateBookVisible(false)
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
                const newEdges = fetchMoreResult ? fetchMoreResult.books.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.books.totalCount : 0
                setPagination({
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                })
                return newEdges.length
                    ? {
                        books: {
                            __typename: previousResult.books.__typename,
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
        { title: 'isbn', dataIndex: 'isbn', key: 'isbn', width: 100 },
        { title: '书名', dataIndex: 'name', key: 'name', width: 150 },
        {
            title: '封面', dataIndex: 'cover', key: 'cover', width: 180,
            render: (value: string) => <Img src={value ? value : ''} />
        },
        {
            title: '作者', dataIndex: 'author', key: 'author', width: 210,
            render: (values: string[], record: any) => {
                if (values) {
                    const tagNodes = values.map((value: string, index: number) => {
                        return (
                            <Tag color={'cyan'} key={"author_" + record.id + "_" + index}>
                                {value}
                            </Tag>
                        )
                    })
                    return <div>{tagNodes}</div>
                }
                return <div />
            }
        },
        {
            title: '译者', dataIndex: 'translator', key: 'translator', width: 100,
            render: (values: string[], record: any) => {
                if (values) {
                    const tagNodes = values.map((value: string, index: number) => {
                        return (
                            <Tag color={'cyan'} key={"translator_" + record.id + "_" + index}>
                                {value}
                            </Tag>
                        )
                    })
                    return <div>{tagNodes}</div>
                }
                return <div />
            }
        },
        { title: '出版社', dataIndex: 'publishingHouse', key: 'publishingHouse', width: 100, },
        { title: '包装', dataIndex: 'packing', key: 'publishingHouse', width: 100, },
        { title: '页数', dataIndex: 'pageSize', key: 'publishingHouse', width: 100, },

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
                        setUpdateBookID(record.id)
                        setUpdateBookVisible(true)
                    }}>编辑</Button><Button
                        onClick={() => {
                            history.push(AdminPath.BOOK_DETAIL + "?id=" + record.id)
                        }}>详情</Button>
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
                新增书籍
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <BookCreateForm
                visible={visible}
                onCreate={onBookCreate}
                onCancel={() => {
                    setVisible(false)
                }}
            />
            <BookUpdateForm
                visible={updateBookVisible}
                id={updateBookID}
                onUpdate={onBookUpdate}
                onCancel={() => {
                    setUpdateBookVisible(false)
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
                    total: data ? data.books.totalCount : 0
                }}
                dataSource={data ? data.books.edges : []}
            />
        </div>

    )
}
