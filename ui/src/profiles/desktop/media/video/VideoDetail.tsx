import React, { useEffect, useState } from "react"
import { message } from "antd"
import "../../../../style/button.less"
import { useQuery } from "@apollo/react-hooks";
import { GET_VIDEO } from '../../../../consts/video.gql';
import Img from "../../../../components/img";
import { VideoPlayer } from "../../../../components/VideoPlayer";
import { useRouteMatch } from "react-router-dom";
import TextArea from "antd/lib/input/TextArea";
import VideoSelect from "../../../common/media/VideoSelect";
import VideoSeriesSelect from "../../../common/media/VideoSeriesSelect";

export default function VideoDetail() {
    const match = useRouteMatch('/app/media/videos/:id');
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
            },
            fetchPolicy: "cache-and-network"
        })
    const [num, setNum] = useState(0)

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
            }
        }

    }

    useEffect(() => {
        if (data && data.videos.edges) {
            const videos = data.videos.edges
            const video = videos.length > 0 ? videos[0] : null
            if (video && video.episodes && video.episodes.length > 0) {
                let episodeNumMap = new Map<number, number>()
                video.episodes.map((value: any, index: number) => {
                    episodeNumMap.set(value.id, index)
                    return value
                })
                setNum(episodeNumMap.get(data.historyInfo?.subSourceID) || 0)
            }
        }
    }, [data])

    useEffect(() => {
        if (video && video.episodes && num < video.episodes.length) {
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
                    videoID={videoItem.id}
                    episodeID={episodeItem.id}
                    url={episodeItem.url}
                    subtitles={episodeItem.subtitles}
                    height={"100%"}
                    width={"100%"}
                    autoplay={false}
                    autoDestroy={false}
                    currentTime={(Number(data?.historyInfo?.subSourceID) !== 0 && Number(data?.historyInfo?.subSourceID) === Number(episodeItem.id)) ? data?.historyInfo?.currentTime : 0}
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
                <VideoSelect data={video?.episodes} num={num} setNum={setNum} />
                <br />
                <VideoSeriesSelect
                    data={data?.videoSerieses.edges}
                    videoID={Number(params.id)} />
            </div>
        </div>)
}
