import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import { VIDEO_RANDOM_TAG_SUGGEST } from 'src/consts/video.gql'
import {
    SyncOutlined,
} from '@ant-design/icons';
import "src/styles/video.less"
import { isMobile } from 'src/utils/util'
import { MobileVideoCard, VideoCard } from './VideoCard'

interface IVideoTagSuggestProps {
    tag: string
    pageSize: number
    color?: string
    fontSize?: number
}

const VideoTagSuggest: React.FC<IVideoTagSuggestProps> = ({
    tag,
    pageSize,
    color = "#85dbf5",
    fontSize = 12
}) => {
    const [loading, setLoading] = useState(false)
    const { error, data, refetch } = useQuery(VIDEO_RANDOM_TAG_SUGGEST,
        {
            variables: {
                tag: tag,
                pageSize: pageSize
            }
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const videos = data ? data.searchVideo.edges : []
    const cards = useMemo(() => {
        return videos.map((item: any, index: number) => {
            if (isMobile()) {
                return <MobileVideoCard
                    key={index}
                    videoID={item.id}
                    cover={item.cover}
                    title={item.title}
                    totalNum={item.totalNum}
                />
            }
            return <VideoCard
                key={index}
                videoID={item.id}
                cover={item.cover}
                title={item.title}
                totalNum={item.totalNum}
            />
        })
    }, [videos])

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
        <div style={{ display: "flex", flexDirection: "column" }}>
            <div style={{ fontSize: fontSize + 6, paddingLeft: 20, fontWeight: 800, textAlign: "left" }}>{tag}</div>
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
