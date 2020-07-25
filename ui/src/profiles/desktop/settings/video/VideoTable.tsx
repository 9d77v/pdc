import { Table, Button, message, Tag, Modal } from 'antd';
import React, { useState, useEffect } from 'react'

import { LIST_VIDEO, ADD_VIDEO, UPDATE_VIDEO, ADD_EPISODE, UPDATE_EPISODE, UPDATE_SUBTITLE, UPDATE_MOBILE_VIDEO } from '../../../../consts/video.gql';
import { useQuery } from '@apollo/react-hooks';
import { VideoCreateForm } from './VideoCreateForm';
import { useMutation } from '@apollo/react-hooks';
import { EpisodeCreateForm } from './EpisodeCreateForm';
import moment from 'moment';
import { VideoPlayer } from '../../../../components/VideoPlayer';
import { VideoUpdateForm } from './VideoUpdateForm';
import { EpisodeUpdateForm } from './EpisodeUpdateForm';
import { Img } from '../../../../components/Img';
import TextArea from 'antd/lib/input/TextArea';
import { TablePaginationConfig } from 'antd/lib/table';
import { PlaySquareTwoTone } from '@ant-design/icons';
import { SubtitleUpdateForm } from './SubtitleUpdateForm';
import Search from 'antd/lib/input/Search';
import { MobileVideoUpdateForm } from './MobileVideoUpdateForm';


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
                            episodeID: record.id,
                            title: episodeRawData.title + " 第" + record.num + "话",
                            url: value,
                            subtitles: record.subtitles,
                            visible: true
                        })}
                        icon={<PlaySquareTwoTone />}
                    />
                )
            }
        },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作',
            dataIndex: 'operation',
            key: 'operation',
            render: (text: any, record: any) => <span>
                <Button
                    onClick={() => {
                        setUpdateEpisodeData({
                            id: record.id,
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
    ];

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
    const [visible, setVisible] = useState(false);
    const [currentVideoID, setCurrentVideoID] = useState(0);
    const [episodeVisible, setEpisodeVisible] = useState(false);
    const [updateVideoVisible, setUpdateVideoVisible] = useState(false)
    const [updateVideoData, setUpdateVideoData] = useState({
        title: "",
        desc: "",
        cover: "",
        pubDate: 0,
        tags: [],
        isShow: false,
    })
    const [updateEpisodeVisible, setUpdateEpisodeVisible] = useState(false)
    const [updateEpisodeData, setUpdateEpisodeData] = useState({
        id: 0,
        num: 0,
        title: "",
        desc: "",
        cover: "",
        url: "",
        subtitles: [],
    })
    const [updateSubtitleVisible, setUpdateSubtitleVisible] = useState(false)
    const [updateMobileVideoVisible, setUpdateMobileVideoVisible] = useState(false)

    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [playerData, setPlayerData] = useState({
        episodeID: 0,
        title: "",
        url: "",
        subtitles: null,
        visible: false
    })
    const [keyword, setKeyword] = useState("")
    const [addVideo] = useMutation(ADD_VIDEO);
    const [updateVideo] = useMutation(UPDATE_VIDEO)
    const [addEpisode] = useMutation(ADD_EPISODE)
    const [updateEpisode] = useMutation(UPDATE_EPISODE)
    const [updateSubtitle] = useMutation(UPDATE_SUBTITLE)
    const [updateMobileVideo] = useMutation(UPDATE_MOBILE_VIDEO)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_VIDEO,
        {
            variables: {
                page: pagination.current,
                pageSize: pagination.pageSize,
                keyword: keyword,
                sorts: [{
                    field: 'id',
                    isAsc: false
                }]
            },
            fetchPolicy: "cache-and-network"
        })
    const [num, setNum] = useState(1);
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const onVideoCreate = async (values: any) => {
        let subtitles = undefined
        if (values.subtitles && values.subtitles.length > 0) {
            subtitles = {
                "name": values.subtitle_lang,
                "urls": values.subtitles
            }
        }
        await addVideo({
            variables: {
                "input": {
                    "title": values.title,
                    "desc": values.desc,
                    "cover": values.cover,
                    "pubDate": values.pubDate ? values.pubDate.unix() : 0,
                    "tags": values.tags || [],
                    "isShow": values.isShow,
                    "videoURLs": values.videoURLs,
                    "subtitles": subtitles
                }
            }
        });
        setVisible(false);
        await refetch()
    };

    const onVideoUpdate = async (values: any) => {
        await updateVideo({
            variables: {
                "input": {
                    "id": currentVideoID,
                    "title": values.title,
                    "desc": values.desc,
                    "cover": values.cover,
                    "pubDate": values.pubDate ? values.pubDate.unix() : 0,
                    "tags": values.tags || [],
                    "isShow": values.isShow,
                }
            }
        });
        setUpdateVideoVisible(false);
        await refetch()
    };

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
        });
        setEpisodeVisible(false);
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
        });
        setUpdateEpisodeVisible(false);
        await refetch()
    };

    const onSubtitleUpdate = async (values: any) => {
        await updateSubtitle({
            variables: {
                "input": {
                    "id": currentVideoID,
                    "subtitles": {
                        "name": values.subtitle_lang,
                        "urls": values.subtitles
                    }
                }
            }
        });
        setUpdateSubtitleVisible(false);
        await refetch()
    };

    const onMobileVideoUpdate = async (values: any) => {
        await updateMobileVideo({
            variables: {
                "input": {
                    "id": currentVideoID,
                    "videoURLs": values.videoURLs,
                }
            }
        });
        setUpdateMobileVideoVisible(false);
    }

    const onChange = (pageConfig: TablePaginationConfig) => {
        fetchMore({
            variables: {
                page: pageConfig.current || 1,
                pageSize: pageConfig.pageSize || 10
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.videos.edges : [];
                const totalCount = fetchMoreResult ? fetchMoreResult.videos.totalCount : 0;
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
            render: (value: number) => moment(value * 1000).format("YYYY年MM月DD日")
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
                        );
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
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', key: 'operation', render: (value: any, record: any) =>
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
                    <Button
                        onClick={() => {
                            setCurrentVideoID(record.id)
                            setUpdateMobileVideoVisible(true)
                        }}>补充视频</Button>
                </span>
        }
    ];
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setVisible(true)
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
            <VideoCreateForm
                visible={visible}
                onCreate={onVideoCreate}
                onCancel={() => {
                    setVisible(false);
                }}
            />
            <VideoUpdateForm
                visible={updateVideoVisible}
                data={updateVideoData}
                onUpdate={onVideoUpdate}
                onCancel={() => {
                    setUpdateVideoVisible(false);
                }}
            />
            <EpisodeCreateForm
                visible={episodeVisible}
                onCreate={onEpisodeCreate}
                onCancel={() => {
                    setEpisodeVisible(false);
                }}
                num={num}
            />
            <EpisodeUpdateForm
                visible={updateEpisodeVisible}
                data={updateEpisodeData}
                onUpdate={onEpisodeUpdate}
                onCancel={() => {
                    setUpdateEpisodeVisible(false);
                }}
            />
            <SubtitleUpdateForm
                visible={updateSubtitleVisible}
                onUpdate={onSubtitleUpdate}
                onCancel={() => {
                    setUpdateSubtitleVisible(false);
                }}
            />
            <MobileVideoUpdateForm
                visible={updateMobileVideoVisible}
                videoID={currentVideoID}
                onUpdate={onMobileVideoUpdate}
                onCancel={() => {
                    setUpdateMobileVideoVisible(false);
                }}
            />
            <Modal
                visible={playerData.visible}
                title={playerData.title}
                okText="确定"
                destroyOnClose={true}
                cancelText="取消"
                width={1008}
                onCancel={
                    () => {
                        setPlayerData({
                            episodeID: 0,
                            title: "",
                            url: "",
                            subtitles: null,
                            visible: false
                        })
                    }
                }
                getContainer={false}
                onOk={() => {
                    setPlayerData({
                        episodeID: 0,
                        title: "",
                        url: "",
                        subtitles: null,
                        visible: false
                    })
                }}
            >  <VideoPlayer
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

    );
}
