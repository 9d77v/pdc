import { Table, Button, message } from 'antd'
import React, { useState, useEffect } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { VideoSeriesCreateForm } from './VideoSeriesCreateForm'
import { useMutation } from '@apollo/react-hooks'
import { VideoSeriesItemCreateForm } from './VideoSeriesItemCreateForm'
import dayjs from 'dayjs'
import { VideoSeriesUpdateForm } from './VideoSeriesUpdateForm'
import { VideoSeriesItemUpdateForm } from './VideoSeriesItemUpdateForm'
import { TablePaginationConfig } from 'antd/lib/table'
import Search from 'antd/lib/input/Search'
import { ADD_VIDEO_SERIES, ADD_VIDEO_SERIES_ITEM, UPDATE_VIDEO_SERIES, UPDATE_VIDEO_SERIES_ITEM } from 'src/gqls/video/mutation'
import { LIST_VIDEO_SERIES } from 'src/gqls/video/query'


function VideoSeriesItemTable(itemRawData: any, setUpdateItemData: any, setUpdateItemVisible: any) {
    const itemData = itemRawData === undefined ? [] : itemRawData.items
    const columns = [
        { title: '视频id', dataIndex: 'videoID', key: 'videoID' },
        { title: '视频名称', dataIndex: 'title', key: 'title' },
        { title: '别名', dataIndex: 'alias', key: 'alias' },
        {
            title: '操作',
            dataIndex: 'operation',
            key: 'operation',
            width: 120,
            render: (text: any, record: any) => <span>
                <Button
                    onClick={() => {
                        setUpdateItemData({
                            videoID: record.videoID,
                            videoSeriesID: record.videoSeriesID,
                            alias: record.alias,
                        })
                        setUpdateItemVisible(true)
                    }}>编辑视频</Button>
            </span>
        },
    ]

    return <div>
        <Table
            rowKey={record => record.videoSeriesID + "_" + record.videoID}
            columns={columns}
            scroll={{ x: 1300 }}
            dataSource={itemData}
            pagination={false} />
    </div>
}

export default function VideoSeriesTable() {
    const [currentVideoSeriesID, setCurrentVideoSeriesID] = useState(0)

    const [videoSeriesVisible, setVideoSeriesVisible] = useState(false)
    const [videoSeriesItemVisible, setVideoSeriesItemVisible] = useState(false)
    const [updateVideoSeriesVisible, setUpdateVideoSeriesVisible] = useState(false)
    const [updateVideoSeriesData, setUpdateVideoSeriesData] = useState({
        id: 0,
        name: ""
    })
    const [updateVideoSeriesItemVisible, setUpdateVideoSeriesItemVisible] = useState(false)
    const [updateVideoSeriesItemData, setUpdateVideoSeriesItemData] = useState({
        videoID: 0,
        videoSeriesID: 0,
        alias: "",
    })

    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })

    const [keyword, setKeyword] = useState("")
    const [addVideoSeries] = useMutation(ADD_VIDEO_SERIES)
    const [updateVideoSeries] = useMutation(UPDATE_VIDEO_SERIES)
    const [addVideoSeriesItem] = useMutation(ADD_VIDEO_SERIES_ITEM)
    const [updateVideoSeriesItem] = useMutation(UPDATE_VIDEO_SERIES_ITEM)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_VIDEO_SERIES,
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
                }
            },
            fetchPolicy: "cache-and-network"
        })
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const onVideoSeriesCreate = async (values: any) => {
        await addVideoSeries({
            variables: {
                "input": {
                    "name": values.name,
                }
            }
        })
        setVideoSeriesVisible(false)
        await refetch()
    }

    const onVideoSeriesUpdate = async (values: any) => {
        await updateVideoSeries({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name
                }
            }
        })
        setUpdateVideoSeriesVisible(false)
        await refetch()
    }

    const onVideoSeriesItemCreate = async (values: any) => {
        await addVideoSeriesItem({
            variables: {
                "input": {
                    "videoID": values.videoID,
                    "videoSeriesID": values.videoSeriesID,
                    "alias": values.alias
                }
            }
        })
        setVideoSeriesItemVisible(false)
        await refetch()
    }

    const onVideoSeriesItemUpdate = async (values: any) => {
        await updateVideoSeriesItem({
            variables: {
                "input": {
                    "videoID": values.videoID,
                    "videoSeriesID": values.videoSeriesID,
                    "alias": values.alias,
                }
            }
        })
        setUpdateVideoSeriesItemVisible(false)
        await refetch()
    }

    const onChange = (pageConfig: TablePaginationConfig) => {
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
                const newEdges = fetchMoreResult ? fetchMoreResult.videoSerieses.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.videoSerieses.totalCount : 0
                const t = {
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                }
                setPagination(t)
                return newEdges.length
                    ? {
                        videoSerieses: {
                            __typename: previousResult.videoSerieses.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult
            }
        })
    }

    const columns = [
        { title: 'ID', dataIndex: 'id', key: 'id' },
        { title: '名称', dataIndex: 'name', key: 'name', width: 180 },

        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', key: 'operation', width: 120, render: (value: any, record: any) =>
                <span><Button
                    onClick={() => {
                        setUpdateVideoSeriesData({
                            id: record.id,
                            name: record.name,
                        })
                        setUpdateVideoSeriesVisible(true)
                    }}>编辑视频系列</Button>
                    <Button
                        onClick={() => {
                            setCurrentVideoSeriesID(record.id)
                            setVideoSeriesItemVisible(true)
                        }}>新增视频</Button>
                </span>
        }
    ]
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setVideoSeriesVisible(true)
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 120 }}
            >
                新增视频系列
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <VideoSeriesCreateForm
                visible={videoSeriesVisible}
                onCreate={onVideoSeriesCreate}
                onCancel={() => {
                    setVideoSeriesVisible(false)
                }}
            />
            <VideoSeriesUpdateForm
                visible={updateVideoSeriesVisible}
                data={updateVideoSeriesData}
                onUpdate={onVideoSeriesUpdate}
                onCancel={() => {
                    setUpdateVideoSeriesVisible(false)
                }}
            />
            <VideoSeriesItemCreateForm
                visible={videoSeriesItemVisible}
                onCreate={onVideoSeriesItemCreate}
                onCancel={() => {
                    setVideoSeriesItemVisible(false)
                }}
                video_series_id={currentVideoSeriesID}
            />
            <VideoSeriesItemUpdateForm
                visible={updateVideoSeriesItemVisible}
                data={updateVideoSeriesItemData}
                onUpdate={onVideoSeriesItemUpdate}
                onCancel={() => {
                    setUpdateVideoSeriesItemVisible(false)
                }}
            />
            <Table
                loading={loading}
                rowKey={record => record.id}
                className="components-table-demo-nested"
                columns={columns}
                expandable={{
                    expandedRowRender: (record: any) => {
                        return VideoSeriesItemTable(record, setUpdateVideoSeriesItemData, setUpdateVideoSeriesItemVisible)
                    }
                }}
                scroll={{ x: 1300 }}
                onChange={onChange}
                pagination={{
                    ...pagination,
                    total: data ? data.videoSerieses.totalCount : 0
                }}
                dataSource={data ? data.videoSerieses.edges : []}
            />
        </div>

    )
}
