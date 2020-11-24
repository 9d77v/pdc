import React, { useEffect, useState } from "react"
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import { GET_VIDEO } from 'src/consts/video.gql'
import { VideoPlayer } from "src/components/videoplayer"
import { useHistory, useLocation } from "react-router-dom"
import { NavBar, Icon } from "antd-mobile"
import VideoSelect from "src/profiles/common/video/VideoSelect"
import VideoSeriesSelect from "src/profiles/common/video/VideoSeriesSelect"
import SimilarVideoList from "src/profiles/common/video/SimilarVideoList"

export default function VideoDetail() {
    const history = useHistory()
    const [episodeItem, setEpisodeItem] = useState({
        id: 0,
        url: "",
        subtitles: null
    })
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const videoID = query.get("video_id")
    const { error, data } = useQuery(GET_VIDEO,
        {
            variables: {
                videoID: videoID
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
        episodes: [],
        theme: ""
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
                    episodes: video.episodes,
                    theme: video.theme
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
                    currentTime={(Number(data?.historyInfo?.subSourceID) !== 0 && Number(data?.historyInfo?.subSourceID) === Number(episodeItem.id)) ? data?.historyInfo?.currentTime : 0}
                />
            </div>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                height: "calc(100% - 276px)", overflowY: "scroll"
            }}>
                <div style={{ display: "flex", flexDirection: 'column' }}>
                    <VideoSelect
                        data={video?.episodes}
                        num={num}
                        setNum={setNum} />
                    <br />
                    <VideoSeriesSelect
                        data={data?.videoSerieses.edges}
                        videoID={Number(videoID)} />
                    <SimilarVideoList videoID={Number(videoID)} pageSize={10} />
                </div>
            </div>
        </div>)
}
