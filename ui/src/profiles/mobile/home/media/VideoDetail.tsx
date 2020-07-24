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
                        return <div key={"pdc-button-" + value.id} className={"pdc-button-selected"}  >{value.num}</div>
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
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                onLeftClick={() => history.push("/app/media/videos")}
            >{videoItem.title + " 第" + (num + 1) + "话"} </NavBar>
            <VideoPlayer
                episodeID={episodeItem.id}
                url={episodeItem.url}
                subtitles={episodeItem.subtitles}
                height={231}
                width={"100%"}
                autoplay={true}
                autoDestroy={false}
            />
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                width: "100%", height: "100%"
            }}>
                <div style={{ marginTop: 20, display: "flex", flexDirection: 'column' }}>
                    <span style={{ textAlign: 'left', paddingLeft: 10, marginBottom: 10 }}> 选集</span>
                    <div>{buttons}</div>
                </div>
            </div>
        </div>)
}
