import { Table, Button, message, Tag, Modal } from 'antd'
import React, { useState, useEffect } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { useMutation } from '@apollo/react-hooks'
import { EpisodeCreateForm } from './EpisodeCreateForm'
import dayjs from 'dayjs'
import { VideoPlayer } from 'src/components/videoplayer'
import { VideoUpdateForm } from './VideoUpdateForm'
import { EpisodeUpdateForm } from './EpisodeUpdateForm'
import Img from 'src/components/img'
import TextArea from 'antd/lib/input/TextArea'
import { TablePaginationConfig } from 'antd/lib/table'
import { PlaySquareTwoTone } from '@ant-design/icons'
import { SubtitleUpdateForm } from './SubtitleUpdateForm'
import Search from 'antd/lib/input/Search'
import { useHistory } from 'react-router-dom'
import { AdminPath } from 'src/consts/path'
import { ADD_EPISODE, SAVE_SUBTITLES, UPDATE_EPISODE, UPDATE_VIDEO } from 'src/gqls/video/mutation'
import { LIST_VIDEO } from 'src/gqls/video/query'


function EpisodeTable(episodeRawData: any, setUpdateEpisodeData: any, setUpdateEpisodeVisible: any, setPlayerData: any) {
    const episodeData = episodeRawData === undefined ? [] : episodeRawData.episodes
    const columns = [
        { title: 'EpisodeID', dataIndex: 'id', key: 'id' },
        { title: '话', dataIndex: 'num', key: 'num' },
        { title: '标题', dataIndex: 'title', key: 'title' },
        {
            title: '简介', dataIndex: 'desc', key: 'desc',
            ellipsis: true
        },
        {
            title: '封面', dataIndex: 'cover', key: 'cover',
            render: (value: string) => <Img src={value} />
        }, {
            title: '视频', dataIndex: 'url', key: 'url',
            render: (value: string, record: any) => {
                return (
                    <Button
                        onClick={() => setPlayerData({
                            videoID: episodeRawData.id,
                            episodeID: record.id,
                            title: episodeRawData.title + " 第" + record.num + "话",
                            url: value,
                            subtitles: record.subtitles,
                            visible: true,
                            theme: episodeRawData.theme
                        })}
                        icon={<PlaySquareTwoTone />}
                    />
                )
            }
        },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作',
            dataIndex: 'operation',
            key: 'operation',
            width: 120,
            render: (text: any, record: any) => <span>
                <Button
                    onClick={() => {
                        setUpdateEpisodeData({
                            id: record.id,
                            videoID: episodeRawData.id,
                            num: record.num,
                            title: record.title,
                            desc: record.desc,
                            cover: record.cover,
                            url: record.url,
                            subtitles: record.subtitles,
                        })
                        setUpdateEpisodeVisible(true)
                    }}>编辑分集</Button>
            </span>
        },
    ]

    return <div>
        <Table
            rowKey={record => record.id}
            columns={columns}
            scroll={{ x: 1300 }}
            dataSource={episodeData}
            pagination={false} />
    </div>
}

export default function VideoTable() {
    const history = useHistory()
    const [currentVideoID, setCurrentVideoID] = useState(0)
    const [episodeVisible, setEpisodeVisible] = useState(false)
    const [updateVideoVisible, setUpdateVideoVisible] = useState(false)
    const [updateVideoData, setUpdateVideoData] = useState({
        title: "",
        desc: "",
        cover: "",
        pubDate: 0,
        tags: [],
        isShow: false,
        isHideOnMobile: false,
        theme: ""
    })
    const [updateEpisodeVisible, setUpdateEpisodeVisible] = useState(false)
    const [updateEpisodeData, setUpdateEpisodeData] = useState({
        id: 0,
        videoID: 0,
        num: 0,
        title: "",
        desc: "",
        cover: "",
        url: "",
        subtitles: [],
    })
    const [updateSubtitleVisible, setUpdateSubtitleVisible] = useState(false)

    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [playerData, setPlayerData] = useState({
        videoID: 0,
        episodeID: 0,
        title: "",
        url: "",
        subtitles: null,
        visible: false,
        theme: ""
    })
    const [keyword, setKeyword] = useState("")
    const [updateVideo] = useMutation(UPDATE_VIDEO)
    const [addEpisode] = useMutation(ADD_EPISODE)
    const [updateEpisode] = useMutation(UPDATE_EPISODE)
    const [saveSubtitles] = useMutation(SAVE_SUBTITLES)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_VIDEO,
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
    const [num, setNum] = useState(1)
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const onVideoUpdate = async (values: any) => {
        await updateVideo({
            variables: {
                "input": {
                    "id": currentVideoID,
                    "title": values.title,
                    "desc": values.desc,
                    "cover": values.cover === "" ? undefined : values.cover,
                    "pubDate": values.pubDate ? values.pubDate.unix() : 0,
                    "tags": values.tags || [],
                    "isShow": values.isShow,
                    "isHideOnMobile": values.isHideOnMobile,
                    "theme": values.theme
                }
            }
        })
        setUpdateVideoVisible(false)
        await refetch()
    }

    const onEpisodeCreate = async (values: any) => {
        await addEpisode({
            variables: {
                "input": {
                    "videoID": currentVideoID,
                    "num": values.num,
                    "title": values.title,
                    "desc": values.desc,
                    "url": values.url,
                    "subtitles": values.subtitles
                }
            }
        })
        setEpisodeVisible(false)
        await refetch()
    }

    const onEpisodeUpdate = async (values: any) => {
        await updateEpisode({
            variables: {
                "input": {
                    "id": values.id,
                    "num": values.num,
                    "title": values.title,
                    "desc": values.desc,
                    "cover": values.cover,
                    "url": values.url,
                    "subtitles": values.subtitles
                }
            }
        })
        setUpdateEpisodeVisible(false)
        await refetch()
    }

    const onSubtitleUpdate = async (values: any) => {
        await saveSubtitles({
            variables: {
                "input": {
                    "id": currentVideoID,
                    "subtitles": {
                        "name": values.subtitle_lang,
                        "urls": values.subtitles
                    }
                }
            }
        })
        setUpdateSubtitleVisible(false)
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
                const newEdges = fetchMoreResult ? fetchMoreResult.videos.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.videos.totalCount : 0
                const t = {
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                }
                setPagination(t)
                return newEdges.length
                    ? {
                        videos: {
                            __typename: previousResult.videos.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult
            }
        })
    }
    const getNum = (currentVideoID: number) => {
        const mediaMap = new Map<number, number>()
        for (const v of data ? data.videos.edges : []) {
            const episodeData = v.episodes
            if (episodeData.length > 0) {
                mediaMap.set(v.id, episodeData[episodeData.length - 1].num + 1)
            } else {
                mediaMap.set(v.id, 1)
            }
        }
        let num = 1
        if (currentVideoID > 0) {
            const i = mediaMap.get(currentVideoID)
            num = i === undefined ? 1 : i
        }
        return num
    }

    const columns = [
        { title: 'ID', dataIndex: 'id', key: 'id' },
        { title: '标题', dataIndex: 'title', key: 'title', width: 180 },
        {
            title: '简介', dataIndex: 'desc', key: 'desc', width: 300,
            render: (value: string) =>
                <TextArea
                    value={value}
                    rows={4}
                    contentEditable={false}
                    style={{
                        backgroundColor: 'rgba(255, 255, 255, 0)',
                        border: 0,
                    }} />
        },
        {
            title: '封面', dataIndex: 'cover', key: 'cover',
            render: (value: string) => <Img src={value} height={107} width={80} />
        },
        {
            title: '上映时间', dataIndex: 'pubDate', key: 'pubDate',
            render: (value: number) => dayjs(value * 1000).format("YYYY年MM月DD日")
        },
        {
            title: '总话数', dataIndex: 'total', key: 'total',
            render: (value: number, record: any) => record.episodes.length
        }, {
            title: '标签', dataIndex: 'tags', key: 'tags',
            render: (values: string[], record: any) => {
                if (values) {
                    const tagNodes = values.map((value: string, index: number) => {
                        return (
                            <Tag color={'cyan'} key={"tag_" + record.id + "_" + index}>
                                {value}
                            </Tag>
                        )
                    })
                    return <div>{tagNodes}</div>
                }
                return <div />
            }
        }, {
            title: '是否显示', dataIndex: 'isShow', key: 'isShow',
            render: (value: Boolean, record: any) => (
                value ? "是" : "否"
            )
        },
        {
            title: '是否手机隐藏', dataIndex: 'isHideOnMobile', key: 'isHideOnMobile',
            render: (value: Boolean, record: any) => (
                value ? "是" : "否"
            )
        },
        // {
        //     title: '主题', dataIndex: 'theme', key: 'theme', width: 80,
        //     render: (value: string, record: any) => (
        //         value === "vjs-theme-lemon" ? "柠檬" : "默认"
        //     )
        // },
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
                        setCurrentVideoID(record.id)
                        setUpdateVideoData({
                            title: record.title,
                            desc: record.desc,
                            cover: record.cover,
                            pubDate: record.pubDate,
                            tags: record.tags || [],
                            isShow: record.isShow,
                            isHideOnMobile: record.isHideOnMobile,
                            theme: record.theme
                        })
                        setUpdateVideoVisible(true)
                    }}>编辑视频</Button>
                    <Button
                        onClick={() => {
                            setCurrentVideoID(record.id)
                            setNum(getNum(record.id))
                            setEpisodeVisible(true)
                        }}>新增分集</Button>
                    <Button
                        onClick={() => {
                            setCurrentVideoID(record.id)
                            setUpdateSubtitleVisible(true)
                        }}>更换字幕</Button>
                </span>
        }
    ]
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    history.push(AdminPath.VIDEO_CREATE)
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 100 }}
            >
                新增视频
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <VideoUpdateForm
                visible={updateVideoVisible}
                data={updateVideoData}
                onUpdate={onVideoUpdate}
                onCancel={() => {
                    setUpdateVideoVisible(false)
                }}
            />
            <EpisodeCreateForm
                currentVideoID={currentVideoID}
                visible={episodeVisible}
                onCreate={onEpisodeCreate}
                onCancel={() => {
                    setEpisodeVisible(false)
                }}
                num={num}
            />
            <EpisodeUpdateForm
                visible={updateEpisodeVisible}
                data={updateEpisodeData}
                onUpdate={onEpisodeUpdate}
                onCancel={() => {
                    setUpdateEpisodeVisible(false)
                }}
            />
            <SubtitleUpdateForm
                visible={updateSubtitleVisible}
                onUpdate={onSubtitleUpdate}
                onCancel={() => {
                    setUpdateSubtitleVisible(false)
                }}
            />
            <Modal
                visible={playerData.visible}
                title={playerData.title}
                footer={null}
                destroyOnClose={true}
                width={1008}
                getContainer={false}
                onCancel={
                    () => {
                        setPlayerData({
                            videoID: 0,
                            episodeID: 0,
                            title: "",
                            url: "",
                            subtitles: null,
                            visible: false,
                            theme: ""
                        })
                    }
                }
            >  <VideoPlayer
                    theme={playerData.theme}
                    videoID={playerData.videoID}
                    episodeID={playerData.episodeID}
                    url={playerData.url}
                    subtitles={playerData.subtitles}
                    height={540}
                    width={960}
                /></Modal>
            <Table
                loading={loading}
                rowKey={record => record.id}
                className="components-table-demo-nested"
                columns={columns}
                expandable={{
                    expandedRowRender: (record: any) => {
                        return EpisodeTable(record, setUpdateEpisodeData, setUpdateEpisodeVisible, setPlayerData)
                    }
                }}
                scroll={{ x: 1300 }}
                onChange={onChange}
                pagination={{
                    ...pagination,
                    total: data ? data.videos.totalCount : 0
                }}
                dataSource={data ? data.videos.edges : []}
            />
        </div>
    )
}
