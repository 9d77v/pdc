import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import React, { useEffect, useMemo } from 'react'
import { GET_VIDEO_TAGS } from 'src/consts/video.gql'
import VideoTagSuggest from 'src/profiles/common/video/VideoTagSuggest'

const VideoIndex = () => {
    const { error, data } = useQuery(GET_VIDEO_TAGS,
        {
            variables: {}
        }
    )
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const cards = useMemo(() => {
        if (data && data.searchVideo.aggResults) {
            const tags = data.searchVideo.aggResults
            return tags.slice(0, 12).map((tag: any, index: number) => {
                return <VideoTagSuggest
                    key={index}
                    tag={tag.key}
                    pageSize={3}
                />
            })
        }
    }, [data])
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            marginTop: 45
        }}>
            {cards}
        </div>
    )
}

export default VideoIndex