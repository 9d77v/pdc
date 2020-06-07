import { Table, Button, message } from 'antd';
import React, { useState, useEffect } from 'react'

import { LIST_VIDEO, ADD_VIDEO, UPDATE_VIDEO, ADD_EPISODE, UPDATE_EPISODE } from '../../../consts/video.gql';
import { useQuery } from '@apollo/react-hooks';
import { VideoCreateForm } from './VideoCreateForm';
import { useMutation } from '@apollo/react-hooks';
import { EpisodeCreateForm } from './EpisodeCreateForm';
import moment from 'moment';
import { VideoPlayer } from '../../../components/VideoPlayer';
import { VideoUpdateForm } from './VideoUpdateForm';
import { EpisodeUpdateForm } from './EpisodeUpdateForm';
import { Img } from '../../../components/Img';
function EpisodeTable(episodeRawData: any, setUpdateEpisodeData: any, setUpdateEpisodeVisible: any) {
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
            title: '视频', dataIndex: 'url', key: 'url', width: 490,
            render: (value: string, record: any) => {
                return (
                    <VideoPlayer
                        episodeID={record.id}
                        url={value}
                        subtitles={record.subtitles}
                        height={270}
                        width={480}
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
        // tags: values.tags,
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
    const [addVideo] = useMutation(ADD_VIDEO);
    const [updateVideo] = useMutation(UPDATE_VIDEO)
    const [addEpisode] = useMutation(ADD_EPISODE)
    const [updateEpisode] = useMutation(UPDATE_EPISODE)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_VIDEO,
        {
            variables: {
                page: 1,
                pageSize: 10,
                sorts: [{
                    field: 'id',
                    isAsc: false
                }]
            }
        });
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
                    // "tags": values.tags,
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
                    // "tags": values.tags,
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

    const onChange = async (page: number) => {
        fetchMore({
            variables: {
                page: page
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.Videos.edges : [];
                const totalCount = fetchMoreResult ? fetchMoreResult.Videos.totalCount : 0;
                return newEdges.length
                    ? {
                        Videos: {
                            __typename: previousResult.Videos.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult;
            }
        })
    }
    const getNum = (currentVideoID: number) => {
        const mediaMap = new Map<number, number>()
        for (const v of data ? data.Videos.edges : []) {
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
        { title: '标题', dataIndex: 'title', key: 'title', width: 200 },
        {
            title: '简介', dataIndex: 'desc', key: 'desc', width: 400
        },
        {
            title: '封面', dataIndex: 'cover', key: 'cover',
            render: (value: string) => <Img src={value} />
        },
        {
            title: '上映时间', dataIndex: 'pubDate', key: 'pubDate',
            render: (value: number) => moment(value * 1000).format("YYYY年MM月DD日")
        },
        {
            title: '总话数', dataIndex: 'total', key: 'total',
            render: (value: number, record: any) => record.episodes.length
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
                            // tags: values.tags,
                            isShow: record.isShow,
                        })
                        setUpdateVideoVisible(true)
                    }}>编辑视频</Button>
                    <Button type="dashed"
                        onClick={() => {
                            setCurrentVideoID(record.id)
                            setNum(getNum(record.id))
                            setEpisodeVisible(true)
                        }}>新增分集</Button>
                </span>
        },
    ];
    return (
        <div>
            <Button
                type="primary"
                onClick={() => {
                    setVisible(true);
                }}
                style={{ float: 'left', marginBottom: 12, zIndex: 1 }}
            >
                新增视频
            </Button>
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
            <Table
                loading={loading}
                rowKey={record => record.id}
                className="components-table-demo-nested"
                columns={columns}
                expandable={{
                    expandedRowRender: (record: any) => {
                        return EpisodeTable(record, setUpdateEpisodeData, setUpdateEpisodeVisible)
                    }
                }}
                pagination={{
                    pageSize: 10,
                    onChange: onChange,
                    total: data ? data.Videos.totalCount : 0,
                    locale: 'zh_CN',
                    showQuickJumper: true,
                    hideOnSinglePage: true
                }}
                dataSource={data ? data.Videos.edges : []}
            />
        </div>

    );
}
