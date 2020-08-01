import React, { useEffect, useState } from "react"
import { message } from "antd"
import "./video.less"
import "../../../style/button.less"
import { useQuery } from "@apollo/react-hooks";
import { GET_VIDEO } from '../../../consts/video.gql';
import { Img } from "../../../components/Img";
import { VideoPlayer } from "../../../components/VideoPlayer";
import { useRouteMatch, useHistory } from "react-router-dom";
import TextArea from "antd/lib/input/TextArea";

export default function VideoDetail() {
    const match = useRouteMatch('/app/media/videos/:id');
    const history = useHistory()
    const [num, setNum] = useState(0)
    const [episodeItem, setEpisodeItem] = useState({
        id: 0,
        url: "",
        subtitles: null
    })
    let ids: number[] = []
    let params: any
    if (match) {
        params = match.params
        ids = [params.id]
    }
    const { error, data } = useQuery(GET_VIDEO,
        {
            variables: {
                ids: ids,
                videoID: params.id
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
    let buttons: any = []
    let seriesName: string = ""
    let seriesButtons: any = []
    let video: any = null
    if (data) {
        if (data.videos.edges) {
            const videos = data.videos.edges
            video = videos.length > 0 ? videos[0] : null
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
                        return <div key={"pdc-button-" + value.id} className={"pdc-button"}
                            onClick={() => { setNum(index) }} >{value.num}</div>
                    })
                }
            }
        }
        if (data.videoSerieses.edges && data.videoSerieses.edges.length > 0 && data.videoSerieses.edges[0].items) {
            const items = data.videoSerieses.edges[0].items
            seriesName = data.videoSerieses.edges[0].name
            seriesButtons = items.map((value: any, index: number) => {
                if (Number(params.id) === Number(value.videoID)) {
                    return <div key={"pdc-button-" + value.videoID} className={"pdc-button-selected"} >{value.alias}</div>
                }
                return <div key={"pdc-button-" + value.videoID} className={"pdc-button"}
                    onClick={() => { history.push('/app/media/videos/' + value.videoID) }} >{value.alias}</div>
            })
        }
    }
    useEffect(() => {
        if (video) {
            setEpisodeItem({
                id: video.episodes[num].id,
                url: video.episodes[num].url,
                subtitles: video.episodes[num].subtitles,
            })
        }
    }, [video, num])
    return (
        <div style={{
            display: 'flex', flexDirection: 'row', height: '100%', width: "100%", overflowX: "scroll"
        }}>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                width: 1080, height: 891,
            }}>
                <VideoPlayer
                    episodeID={episodeItem.id}
                    url={episodeItem.url}
                    subtitles={episodeItem.subtitles}
                    height={"100%"}
                    width={"100%"}
                    autoplay={true}
                    autoDestroy={false}
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
                <span style={{ textAlign: 'left', paddingLeft: 10, marginBottom: 10 }}>选集</span>
                <div>{buttons}</div>
                <br />
                <span style={{ textAlign: "left", marginBottom: 10 }}>{seriesName}</span>
                <div>{seriesButtons}</div>
            </div>
            <div style={{ flex: 1 }}>
            </div>
        </div>)
}
