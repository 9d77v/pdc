import { useHistory } from "react-router-dom"
import { AppPath } from "src/consts/path"
import React from 'react'
import Img from "src/components/img"
interface IVideoCardProps {
    videoID: number
    cover: string
    title: string
    totalNum: number
}

export const VideoCard: React.FC<IVideoCardProps> = ({
    videoID,
    cover,
    title,
    totalNum
}) => {
    const history = useHistory()
    return (
        <div
            onClick={() => history.push(AppPath.VIDEO_DETAIL + "?video_id=" + videoID)}
            className={"card"}
        >
            <div style={{ clear: "both" }} />
            <Img src={cover} />
            <div style={{ marginTop: 5, fontSize: 14 }}>{title}</div>
            <div style={{ fontSize: 12 }}>全{totalNum}话</div>
        </div >
    )
}

export const MobileVideoCard: React.FC<IVideoCardProps> = ({
    videoID,
    cover,
    title,
    totalNum
}) => {
    const history = useHistory()
    return (
        <div
            onClick={() => history.push(AppPath.VIDEO_DETAIL + "?video_id=" + videoID)}
            style={{
                width: "30%",
                margin: "2.5% 0 0 2.5%",
                height: 210,
                display: "flex",
                float: "left",
                flexDirection: "column"
            }}
        >
            <Img src={cover} width={"100%"} height={"70%"} />
            <div style={{
                fontSize: 12,
                height: 36,
                overflow: "hidden",
                textOverflow: "ellipsis"
            }}>{title}</div>
            <div style={{ fontSize: 10 }}>全{totalNum}话</div>
        </div >)
}