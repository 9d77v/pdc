import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import React, { FC, useEffect, useMemo, useState } from 'react'
import {
    SyncOutlined,
} from '@ant-design/icons';

import { isMobile } from 'src/utils/util'
import { MobileVideoCard, VideoCard } from './VideoCard'
import { VIDEO_RANDOM_TAG_SUGGEST } from 'src/gqls/video/query';

interface IVideoTagSuggestProps {
    title?: string
    tag: string
    pageSize: number
    color?: string
    fontSize?: number
    width?: number | string
}

const VideoTagSuggest: FC<IVideoTagSuggestProps> = ({
    title,
    width = "100%",
    tag,
    pageSize,
    color = "#85dbf5",
    fontSize = 12
}) => {
    const [loading, setLoading] = useState(false)
    const { error, data, refetch } = useQuery(VIDEO_RANDOM_TAG_SUGGEST,
        {
            variables: {
                searchParam: {
                    pageSize: pageSize,
                    tags: [tag],
                    page: 1,
                    isRandom: true,
                    isMobile: isMobile()
                }
            }
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const cards = useMemo(() => {
        const videos = data ? data.searchVideo.edges : []
        return videos.map((item: any, index: number) => {
            if (isMobile()) {
                return <MobileVideoCard
                    key={index}
                    episodeID={item.episodeID}
                    cover={item.cover}
                    title={item.title}
                    totalNum={item.totalNum}
                />
            }
            return <VideoCard
                key={index}
                episodeID={item.episodeID}
                cover={item.cover}
                title={item.title}
                totalNum={item.totalNum}
            />
        })
    }, [data])

    let timer: any
    const next = () => {
        clearTimeout(timer)
        if (!loading) {
            setLoading(true)
        }
        timer = setTimeout(() => {
            refetch()
            setLoading(false)
        }, 1000)
    }

    return (
        <div style={{ display: "flex", flexDirection: "column", width: width }}>
            <div style={{ fontSize: fontSize + 6, paddingLeft: 20, fontWeight: 800, textAlign: "left" }}>
                {tag ? tag : title}
            </div>
            <div>{cards}</div>
            <div style={{
                fontSize: fontSize,
                padding: 10,
                color: color,
                textAlign: "center",
                cursor: "pointer"
            }}
                onClick={() => next()}>  <SyncOutlined spin={loading} />  换一换</div>
        </div>
    )
}

export default VideoTagSuggest
