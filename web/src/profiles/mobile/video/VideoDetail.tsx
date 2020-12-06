import React, { useEffect, useMemo } from "react"
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import { VideoPlayer } from "src/components/videoplayer"
import { useHistory, useLocation } from "react-router-dom"
import { NavBar, Icon } from "antd-mobile"
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
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >{videoItem.title + " 第" + (num + 1) + "话"} </NavBar>
            <div style={{ marginTop: 45 }}>
                <VideoPlayer
                    theme={videoItem.theme}
                    videoID={videoItem.id}
                    episodeID={episodeItem.id}
                    url={episodeItem.url}
                    subtitles={episodeItem.subtitles}
                    height={231}
                    width={"100%"}
                    autoplay={false}
                    autoDestroy={false}
                    currentTime={(Number(historyInfo?.subSourceID) !== 0 && Number(historyInfo?.subSourceID) === Number(episodeItem.id)) ? historyInfo?.currentTime : 0}
                />
            </div>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                height: "calc(100% - 276px)", overflowY: "scroll"
            }}>
                <div style={{ display: "flex", flexDirection: 'column' }}>
                    <VideoSelect
                        data={video?.episodes}
                        num={num} />
                    <br />
                    <VideoSeriesSelect
                        data={videoSeries}
                        videoID={Number(videoItem.id)} />
                    {Number(videoItem.id) ? <SimilarVideoList videoID={Number(videoItem.id)} pageSize={10} /> : null}
                </div>
            </div>
        </div>)
}
