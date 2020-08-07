import React, { useEffect, useState } from "react"
import { message } from "antd"
import "../../../../style/button.less"
import { useQuery } from "@apollo/react-hooks";
import { GET_VIDEO } from '../../../../consts/video.gql';
import { VideoPlayer } from "../../../../components/VideoPlayer";
import { useRouteMatch, useHistory } from "react-router-dom";
import { NavBar, Icon } from "antd-mobile";

export default function VideoDetail() {
    const match = useRouteMatch('/app/media/videos/:id');
    const history = useHistory()
    const [episodeItem, setEpisodeItem] = useState({
        id: 0,
        url: "",
        mobileURL: "",
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
                            return <div key={"pdc-button-" + value.id} className={"pdc-button-selected"}  >{value.num}</div>
                        }
                        return <div key={"pdc-button-" + value.id} className={"pdc-button"} onClick={() => { setNum(index) }} >{value.num}</div>
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
        if (data && data.historyInfo && data.videos.edges) {
            const videos = data.videos.edges
            const video = videos.length > 0 ? videos[0] : null
            if (video && video.episodes && video.episodes.length > 0) {
                let episodeNumMap = new Map<number, number>()
                video.episodes.map((value: any, index: number) => {
                    episodeNumMap.set(value.id, index)
                    return value
                })
                setNum(episodeNumMap.get(data.historyInfo.subSourceID) || 0)
            }
        }
    }, [data])

    useEffect(() => {
        if (video) {
            setEpisodeItem({
                id: video.episodes[num].id,
                url: video.episodes[num].url,
                mobileURL: video.episodes[num].mobileURL || "",
                subtitles: video.episodes[num].subtitles,
            })
        }
    }, [video, num])

    return (
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                onLeftClick={() => history.push("/app/media/videos")}
            >{videoItem.title + " 第" + (num + 1) + "话"} </NavBar>
            <VideoPlayer
                videoID={videoItem.id}
                episodeID={episodeItem.id}
                url={episodeItem.mobileURL !== "" ? episodeItem.mobileURL : episodeItem.url}
                subtitles={episodeItem.subtitles}
                height={231}
                width={"100%"}
                autoplay={false}
                autoDestroy={false}
                currentTime={(Number(data?.historyInfo?.subSourceID) !== 0 && Number(data?.historyInfo?.subSourceID) === Number(episodeItem.id)) ? data?.historyInfo?.currentTime : 0}
            />
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                width: "100%", height: "100%"
            }}>
                <div style={{ marginTop: 20, display: "flex", flexDirection: 'column' }}>
                    <span style={{ textAlign: 'left', paddingLeft: 10, marginBottom: 10 }}> 选集</span>
                    <div>{buttons}</div>
                    <br />
                    <span style={{ textAlign: "left", marginBottom: 10 }}>{seriesName === "" ? "" : seriesName + "系列"}</span>
                    <div>{seriesButtons}</div>
                </div>
            </div>
        </div>)
}