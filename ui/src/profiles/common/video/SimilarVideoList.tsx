import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import React, { FC, useEffect } from 'react'
import { SIMILAR_VIDEOS } from 'src/consts/video.gql'

interface ISimilarVideoListProps {
    videoID: number
    pageSize: number
}

const SimilarVideoList: FC<ISimilarVideoListProps> = ({
    videoID,
    pageSize
}) => {
    const { error, data } = useQuery(SIMILAR_VIDEOS,
        {
            variables: {
                videoID: videoID,
                pageSize: pageSize
            }
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const videos = data ? data.similarVideos.edges : []
    return (
        <div>
            <div
                style={{ fontSize: 16, fontWeight: 500, textAlign: 'left', paddingLeft: 10 }}>
                {videos.length > 0 ? "相似推荐" : null}
            </div>
        </div>
    )
}

export default SimilarVideoList
