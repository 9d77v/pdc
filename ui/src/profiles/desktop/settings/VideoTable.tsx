import { Table, Button } from 'antd';
import React, { useState } from 'react'

import { LIST_VIDEO, ADD_VIDEO, ADD_EPISODE } from '../../../gqls/video.gql';
import { useQuery } from '@apollo/react-hooks';
import { VideoCreateForm } from './VideoCreateFrom';
import { useMutation } from '@apollo/react-hooks';
import { EpisodeCreateForm } from './EpisodeCreateFrom';
import moment from 'moment';
import { VideoPlayer } from '../../components/VideoPlayer';

function EpisodeTable(record: any) {
    const episodeData = record === undefined ? [] : record.episodes
    const columns = [
        { title: 'EpisodeID', dataIndex: 'id', key: 'id' },
        { title: '话', dataIndex: 'num', key: 'num' },
        { title: '标题', dataIndex: 'title', key: 'title' },
        { title: '简介', dataIndex: 'desc', key: 'desc' },
        { title: '封面', dataIndex: 'cover', key: 'cover' },
        {
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
            // render: (record: any) => <Button onClick={() => {
            //     setEpisodeVisible(true)
            // }}>修改分集</Button>
        },
    ];
    return <Table
        rowKey={record => record.id}
        columns={columns}
        dataSource={episodeData}
        pagination={false} />;
}


export default function VideoTable() {
    const [visible, setVisible] = useState(false);
    const [currentVideoID, setCurrentVideoID] = useState(0);
    const [episodeVisible, setEpisodeVisible] = useState(false);
    const [addVideo] = useMutation(ADD_VIDEO);
    const [addEpisode] = useMutation(ADD_EPISODE)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_VIDEO,
        {
            variables: {
                page: 1,
                pageSize: 10
            }
        });
    const [num, setNum] = useState(1);
    if (error) return <div>Error! ${error}</div>;

    const onVideoCreate = async (values: any) => {
        await addVideo({
            variables: {
                "input": {
                    "title": values.title,
                    "desc": values.desc,
                    "cover": values.cover,
                    "pubDate": values.pubDate ? values.pubDate.unix() : 0,
                    // "tags": values.tags,
                    "isShow": values.isShow,
                }
            }
        });
        setVisible(false);
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

    const onChange = (page: number) => {
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
        { title: '标题', dataIndex: 'title', key: 'title' },
        { title: '简介', dataIndex: 'desc', key: 'desc', width: 400 },
        {
            title: '封面', dataIndex: 'cover', key: 'cover',
            render: (value: string) => <img src={value} width={160} height={210} alt={"图片不存在"} />
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
            title: '操作', key: 'operation', render: (record: any) =>
                <span> <Button onClick={() => {
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
            <EpisodeCreateForm
                visible={episodeVisible}
                onCreate={onEpisodeCreate}
                onCancel={() => {
                    setEpisodeVisible(false);
                }}
                num={num}
            />
            <Table
                loading={loading}
                rowKey={record => record.id}
                className="components-table-demo-nested"
                columns={columns}
                expandable={{
                    expandedRowRender: (record: any) => {
                        return EpisodeTable(record)
                    }
                }}
                pagination={{
                    pageSize: 10,
                    onChange: onChange,
                    total: data ? data.Videos.totalCount : 0,
                    // current: current,
                    locale: 'zh_CN',
                    showQuickJumper: true,
                }}
                dataSource={data ? data.Videos.edges : []}
            />
        </div>

    );
}
