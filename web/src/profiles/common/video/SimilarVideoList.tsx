import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import React, { FC, useEffect, useMemo } from 'react'
import { useHistory } from 'react-router-dom'
import Img from 'src/components/img'
import { AppPath } from 'src/consts/path'
import { SIMILAR_VIDEOS } from 'src/consts/video.gql'
import { VideoCardModel } from 'src/models/video'
import { isMobile } from 'src/utils/util'

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
                input: {
                    videoID: videoID,
                    pageSize: pageSize,
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

    const history = useHistory()
    const videos: JSX.Element[] = useMemo(() => {
        if (data) {
            return data.similarVideos.edges.map((video: VideoCardModel, index: number) => {
                const link = AppPath.VIDEO_DETAIL + "?episode_id=" + video.episodeID
                return (<div key={index}
                    onClick={() => history.push(link)}
                    style={{ display: "flex", cursor: "pointer", paddingBottom: 10 }}>
                    <div style={{ width: 160 }}
                    >
                        <Img
                            src={video.cover}
                            width={160}
                            height={100}
                            hideModal={true} />
                    </div>
                    <div style={{
                        flex: 1,
                        padding: 5,
                        paddingLeft: 10,
                        position: "relative",
                        fontSize: 14,
                        textAlign: "left"
                    }}>
                        {video.title}
                        <div style={{ fontSize: 10, position: "absolute", bottom: 5 }}>
                            全{video.totalNum}话
                            </div>
                    </div>
                </div>)
            })
        }
        return []
    }, [data, history])
    return (
        <div style={{ marginTop: 10 }}>
            <div
                style={{ fontSize: 16, fontWeight: 500, textAlign: 'left', paddingLeft: 10 }}>
                <div style={{ marginBottom: 10 }}>{videos.length > 0 ? "相似推荐" : null}</div>
                {videos}
            </div>
        </div>
    )
}

export default SimilarVideoList
