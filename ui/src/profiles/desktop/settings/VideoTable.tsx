import { Table, Button } from 'antd';
import React, { useState } from 'react'
import { LIST_MEDIA, ADD_MEDIA, ADD_EPISODE } from '../../../stores/videostore';
import { useQuery } from '@apollo/react-hooks';
import { MediaCreateForm } from './MediaCreateFrom';
import { useMutation } from '@apollo/react-hooks';
import { EpisodeCreateForm } from './EpisodeCreateFrom';
import moment from 'moment';


function EpisodeTable(record: any) {
    const columns = [
        { title: 'EpisodeID', dataIndex: 'order', key: 'order' },
        { title: '标题', dataIndex: 'title', key: 'title' },
        { title: '简介', dataIndex: 'desc', key: 'desc' },
        { title: '封面', dataIndex: 'cover', key: 'cover' },
        { title: '视频地址', dataIndex: 'url', key: 'url' },
        { title: '字幕地址', dataIndex: 'subtitle', key: 'subtitle' },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        // {
        //     title: '操作',
        //     dataIndex: 'operation',
        //     key: 'operation'
        // },
    ];
    const episodeData = record === undefined ? [] : record.episodes
    return <Table
        rowKey={record => record.id}
        columns={columns}
        dataSource={episodeData}
        pagination={false} />;
}

export default function VideoTable() {
    const [visible, setVisible] = useState(false);
    const [currentMediaID, setCurrentMediaID] = useState(0);
    const [episodeVisible, setEpisodeVisible] = useState(false);
    const [addMedia] = useMutation(ADD_MEDIA);
    const [addEpisode] = useMutation(ADD_EPISODE)
    const { loading, error, data, refetch } = useQuery(LIST_MEDIA);
    if (error) return <div>Error! ${error}</div>;

    const onMediaCreate = async (values: any) => {
        await addMedia({
            variables: {
                "input": {
                    "title": values.title,
                    "desc": values.desc,
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
                    "mediaID": currentMediaID,
                    "order": values.order.toString(),
                    "title": values.title,
                    "desc": values.desc,
                    "url": values.url,
                }
            }
        });
        setEpisodeVisible(false);
        await refetch()
    };

    const mediaData = data === undefined ? [] : data.listMedia
    const columns = [
        { title: 'ID', dataIndex: 'id', key: 'id' },
        { title: '标题', dataIndex: 'title', key: 'title' },
        { title: '简介', dataIndex: 'desc', key: 'desc' },
        { title: '封面', dataIndex: 'cover', key: 'cover' },
        {
            title: '上映时间', dataIndex: 'pubDate', key: 'pubDate',
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
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
            title: '操作', key: 'operation', render: (record: any) => <Button onClick={() => {
                console.log(record.id)
                setCurrentMediaID(record.id)
                setEpisodeVisible(true)
            }}>新增分集</Button>
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
            <MediaCreateForm
                visible={visible}
                onCreate={onMediaCreate}
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
                dataSource={mediaData}
            />
        </div>

    );
}