import React, { FC, useMemo } from 'react'
import Img from 'src/components/img'
import { AppPath } from 'src/consts/path'
import { VideoCardModel } from 'src/models/video'

interface ISimilarVideoListProps {
    data: any
}

const SimilarVideoList: FC<ISimilarVideoListProps> = ({
    data
}) => {
    const videos: JSX.Element[] = useMemo(() => {
        if (data) {
            return data.map((video: VideoCardModel, index: number) => {
                const link = AppPath.VIDEO_DETAIL + "?epid=" + video.episodeID + "&autoJump=true"
                return (<div key={index}
                    onClick={() => window.location.href = link}
                    style={{ display: "flex", cursor: "pointer", paddingBottom: 10 }}>
                    <div style={{ width: 160 }}
                    >
                        <Img
                            src={video.cover}
                            width={160}
                            height={100}
                        />
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
                </div >)
            })
        }
        return []
    }, [data])
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
