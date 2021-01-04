import React, { FC, useEffect, useMemo } from "react"
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
import { getVideoDetail } from "src/models/video"
import { isMobile } from "src/utils/util"

export const EpisodePage: FC = () => {
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
                    page: 1,
                    pageSize: 50
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
            history.replace(AppPath.VIDEO_DETAIL + "?epid=" + historyInfo.subSourceID)
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
    return (
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >{videoDetail.videoItem.title + " 第" + (videoDetail.num + 1) + "话"} </NavBar>
            <div style={{ marginTop: 45 }}>
                <VideoPlayer
                    theme={videoDetail.videoItem.theme}
                    videoID={videoDetail.videoItem.id}
                    episodeID={videoDetail.episodeItem.id}
                    url={videoDetail.episodeItem.url}
                    subtitles={videoDetail.episodeItem.subtitles}
                    height={231}
                    width={"100%"}
                    autoplay={false}
                    currentTime={(Number(historyInfo?.subSourceID) !== 0 && Number(historyInfo?.subSourceID) === Number(videoDetail.episodeItem.id)) ? historyInfo?.currentTime : 0}
                />
            </div>
            <div style={{
                display: 'flex', flexDirection: 'column', padding: 10,
                height: "calc(100% - 276px)", overflowY: "scroll"
            }}>
                <div style={{ display: "flex", flexDirection: 'column' }}>
                    <VideoSelect
                        data={videoDetail.video?.episodes}
                        num={videoDetail.num} />
                    <br />
                    {videoSeries}
                    {similarVideoList}
                </div>
            </div>
        </div>)
}
