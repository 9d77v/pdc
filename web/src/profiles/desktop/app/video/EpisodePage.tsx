import React, { FC, useEffect, useMemo } from "react"
import { message, Tag } from "antd"
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
import { getVideoDetail } from "src/models/video"
import { isMobile } from "src/utils/util"
import dayjs from "dayjs"

const EpisodePage: FC = () => {
    const history = useHistory()
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const episodeID = query.get("epid")
    const autoJump = query.get("autoJump")

    const { error, data } = useQuery(GET_VIDEO_DETAIL,
        {
            variables: {
                episodeID: episodeID,
                similiarVideoParam: {
                    isMobile: isMobile(),
                    pageSize: 10,
                },
                searchParam: {
                    page: 1
                }
            }
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const historyInfo = data?.histories.edges.length > 0 ? data?.histories.edges[0] : null

    useEffect(() => {
        if (historyInfo && autoJump === "true" && Number(historyInfo.subSourceID !== Number(episodeID))) {
            window.location.replace(AppPath.VIDEO_DETAIL + "?epid=" + historyInfo.subSourceID)
        }
    }, [historyInfo, history, autoJump, episodeID])

    const videoDetail = useMemo(() => {
        return getVideoDetail(data, episodeID)
    }, [data, episodeID])

    const videoSeries = useMemo(() => {
        if (Number(videoDetail.videoItem.id)) {
            return <VideoSeriesSelect data={videoDetail.videoSerieses} videoID={Number(videoDetail.videoItem.id)} />
        }
        return null
    }, [videoDetail])

    const similarVideoList = useMemo(() => {
        if (Number(videoDetail.videoItem.id)) {
            return <SimilarVideoList data={videoDetail.similarVideos} />
        }
        return null
    }, [videoDetail])

    const tagNodes = useMemo(() => {
        const tags = videoDetail.videoItem.tags.map((value: string, index: number) => {
            return (
                <Tag color={'cyan'} key={"tag_" + index}>
                    {value}
                </Tag>
            )
        })
        return <div style={{ marginLeft: 10 }}>{tags}</div>
    }, [videoDetail])

    return (
        <div style={{
            display: 'flex', flexDirection: 'row', height: '100%', width: "100%", overflowX: "scroll"
        }}>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10
            }}>
                <VideoPlayer
                    theme={videoDetail.videoItem.theme}
                    videoID={videoDetail.videoItem.id}
                    episodeID={videoDetail.episodeItem.id}
                    url={videoDetail.episodeItem.url}
                    subtitles={videoDetail.episodeItem.subtitles}
                    height={707}
                    width={1088}
                    autoplay={false}
                    currentTime={(Number(historyInfo?.subSourceID) !== 0 && Number(historyInfo?.subSourceID) === Number(videoDetail.episodeItem.id)) ? historyInfo?.currentTime : 0}
                />
                <div style={{ marginTop: 10, display: 'flex', flexDirection: 'row', flex: 1 }}>
                    <Img src={videoDetail.videoItem.cover} showModal />
                    <div style={{ display: 'flex', textAlign: "left", paddingLeft: 10, flexDirection: 'column', flex: 1, paddingInline: 10 }} >
                        <div style={{ fontSize: 18, display: "flex" }}> {videoDetail.videoItem.title}{tagNodes}</div>
                        <div style={{ paddingRight: 10 }}> 全{videoDetail.videoItem.episodes.length}话</div>
                        <span style={{ marginTop: 15, marginBottom: 10 }}>{dayjs(videoDetail.videoItem.pubDate * 1000).format("YYYY年MM月DD日") + "开播"}</span>
                        <div style={{ paddingRight: 10 }}>
                            <TextArea
                                value={videoDetail.videoItem.desc}
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
            <div style={{ margin: 20, display: "flex", flexDirection: 'column', maxWidth: 350, flex: 1 }}>
                <VideoSelect data={videoDetail.video?.episodes} num={videoDetail.num} />
                <br />
                {videoSeries}
                {similarVideoList}
            </div>
        </div>)
}

export default EpisodePage
