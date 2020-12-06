import React, { useEffect, useMemo } from "react"
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import Img from "src/components/img"
import { VideoPlayer } from "src/components/videoplayer"
import { useHistory, useLocation } from "react-router-dom"
import TextArea from "antd/lib/input/TextArea"
import VideoSelect from "src/profiles/common/video/VideoSelect"
import VideoSeriesSelect from "src/profiles/common/video/VideoSeriesSelect"
import SimilarVideoList from "src/profiles/common/video/SimilarVideoList"
import { AppPath } from "src/consts/path"
import { GET_VIDEO_DETAIL } from "src/gqls/video/query"

export default function VideoDetail() {
    const history = useHistory()
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const episodeID = query.get("episode_id")
    const next = query.get("next")
    const { error, data } = useQuery(GET_VIDEO_DETAIL,
        {
            variables: {
                episodeID: episodeID
            }
        })
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const video = useMemo(() => {
        return data?.videoDetail.video
    }, [data])

    const videoItem = useMemo(() => {
        if (video) {
            return {
                id: video.id,
                cover: video.cover,
                title: video.title,
                desc: video.desc,
                episodes: video.episodes,
                theme: video.theme
            }
        }
        return {
            id: 0,
            cover: "",
            title: "",
            desc: "",
            episodes: [],
            theme: ""
        }
    }, [video])

    const historyInfo = useMemo(() => {
        return data?.videoDetail?.historyInfo
    }, [data])

    const videoSeries = useMemo(() => {
        return data?.videoDetail?.videoSerieses
    }, [data])


    useEffect(() => {
        if (historyInfo && next !== "true" && Number(historyInfo.subSourceID !== Number(episodeID))) {
            history.replace(AppPath.VIDEO_DETAIL + "?episode_id=" + historyInfo.subSourceID)
        }
    }, [historyInfo, history, next, episodeID])

    const num = useMemo(() => {
        if (video && video.episodes) {
            for (let i = 0, len = video.episodes.length; i < len; i++) {
                if (Number(video.episodes[i].id) === Number(episodeID)) {
                    return i
                }
            }
        }
        return 0
    }, [video, episodeID])

    const episodeItem = useMemo(() => {
        if (video && video.episodes && num < video.episodes.length) {
            return {
                id: video.episodes[num].id,
                url: video.episodes[num].url,
                subtitles: video.episodes[num].subtitles,
            }
        }
        return {
            id: 0,
            url: "",
            subtitles: null
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
                    theme={videoItem.theme}
                    videoID={videoItem.id}
                    episodeID={episodeItem.id}
                    url={episodeItem.url}
                    subtitles={episodeItem.subtitles}
                    height={"100%"}
                    width={"100%"}
                    autoplay={false}
                    autoDestroy={false}
                    currentTime={(Number(historyInfo?.subSourceID) !== 0 && Number(historyInfo?.subSourceID) === Number(episodeItem.id)) ? historyInfo?.currentTime : 0}
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
                <VideoSelect data={video?.episodes} num={num} />
                <br />
                <VideoSeriesSelect
                    data={videoSeries}
                    videoID={Number(videoItem.id)} />
                {Number(videoItem.id) ? <SimilarVideoList videoID={Number(videoItem.id)} pageSize={10} /> : null}
            </div>
        </div>)
}
