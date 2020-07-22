import React, { useEffect, useState } from "react"
import { message } from "antd"
import "./video.less"
import "../../../style/button.less"
import { useQuery } from "@apollo/react-hooks";
import { GET_VIDEO } from '../../../consts/video.gql';
import { Img } from "../../../components/Img";
import { VideoPlayer } from "../../../components/VideoPlayer";
import { useRouteMatch } from "react-router-dom";
import TextArea from "antd/lib/input/TextArea";

export default function VideoDetail() {
    const match = useRouteMatch('/app/media/videos/:id');
    const [num, setNum] = useState(0)
    let ids: number[] = []
    if (match) {
        const params: any = match.params
        ids = [params.id]
    }
    const { error, data } = useQuery(GET_VIDEO,
        {
            variables: {
                ids: ids
            }
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    let videoItem = {
        id: 0,
        cover: "",
        title: "",
        desc: "",
        episodes: []
    }
    let episodeItem = {
        id: 0,
        url: "",
        subtitles: null
    }
    let buttons: any = []
    if (data && data.videos.edges) {
        const videos = data.videos.edges
        const video = videos.length > 0 ? videos[0] : null
        if (video) {
            videoItem = ({
                id: video.id,
                cover: video.cover,
                title: video.title,
                desc: video.desc,
                episodes: video.episodes
            })
            if (video.episodes && video.episodes.length > 0) {
                buttons = video.episodes.map((value: any, index: number) => {
                    if (index === num) {
                        return <div key={"pdc-button-" + value.id} className={"pdc-button-selected"} >{value.num}</div>
                    }
                    return <div key={"pdc-button-" + value.id} className={"pdc-button"} onClick={() => { setNum(index) }} >{value.num}</div>
                })
                episodeItem = ({
                    id: video.episodes[num].id,
                    url: video.episodes[num].url,
                    subtitles: video.episodes[num].subtitles
                })
            }
        }
    }

    return (
        <div style={{
            display: 'flex', flexDirection: 'row', height: '100%', width: "100%", overflowX: "scroll"
        }}>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                width: 1080, height: 607.5,
                minWidth: 360, minHeight: 202.5,
            }}>
                <VideoPlayer
                    episodeID={episodeItem.id}
                    url={episodeItem.url}
                    subtitles={episodeItem.subtitles}

                    height={"100%"}
                    width={"100%"}
                    autoplay={true}
                />
                <div style={{ marginTop: 10, display: 'flex', flexDirection: 'row', flex: 1 }}>
                    <Img src={videoItem.cover} />
                    <div style={{ display: 'flex', flexDirection: 'column', flex: 1, paddingInline: 10 }} >
                        <div style={{ textAlign: "left", fontSize: 18, padding: 10 }}> {videoItem.title}</div>
                        <div style={{ textAlign: "left", paddingLeft: 10, paddingRight: 10 }}> 全{videoItem.episodes.length}话</div>
                        <div style={{ textAlign: 'left', paddingLeft: 10, paddingRight: 10 }}>
                            <TextArea
                                value={videoItem.desc}
                                rows={9}
                                contentEditable={false}
                                style={{
                                    backgroundColor: 'rgba(255, 255, 255, 0)',
                                    border: 0,
                                }} />
                        </div>
                    </div>
                </div>
            </div>
            <div style={{ margin: 20, display: "flex", flexDirection: 'column', width: 350 }}>
                <span style={{ textAlign: 'left', paddingLeft: 10, marginBottom: 10 }}> 剧集列表</span>
                <div>{buttons}</div>
            </div>
            <div style={{ flex: 1 }}>
            </div>
        </div>)
}
