import React, { useEffect, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'
import "video.js/dist/video-js.min.css"
import "./index.less"
import "./vjs-theme-lemon.less"

import video_zhCN from 'video.js/dist/lang/zh-CN.json'
import { useLocation } from 'react-router-dom'
import { recordHistory } from 'src/consts/http'

const lang: any = video_zhCN
lang["Picture-in-Picture"] = "画中画"
videojs.addLanguage('zh-CN', lang);
export interface SubtitleProps {
    name: string
    url: string
}

export interface VideoPlayerProps {
    theme?: string
    videoID: number
    episodeID: number
    url: string
    subtitles: [SubtitleProps] | null
    height?: any
    width?: any
    minHeight?: number
    minWidth?: number
    autoplay?: boolean
    autoDestroy?: boolean
    currentTime?: number
}

const VideoPlayer: React.FC<VideoPlayerProps> = ({
    theme = "",
    videoID,
    episodeID,
    url,
    subtitles,
    height,
    width,
    minWidth,
    minHeight,
    autoplay,
    autoDestroy,
    currentTime = 0,
}) => {
    const location = useLocation();
    const isApp = location.pathname.indexOf("/app") >= 0

    const videoKey = "custom-video" + episodeID
    const [videoNode, setVideoNode] = useState()
    const [player, setPlayer] = useState<VideoJsPlayer>()

    const props: VideoJsPlayerOptions = {
        autoplay: autoplay,
        sources: [{
            src: url,
            type: 'video/mp4',
        }],
        language: "zh-CN",
        controls: true,
        playbackRates: [0.5, 1, 2, 4, 16],
        loop: false,
    };

    if (autoDestroy === undefined) {
        autoDestroy = true
    }
    useEffect(() => {
        if (videoNode && url) {
            if (!player) {
                let tmpPlayer = videojs(videoNode, props, () => {
                    for (const item of subtitles || []) {
                        tmpPlayer.addRemoteTextTrack({
                            "kind": "subtitles",
                            "src": item.url,
                            "label": item.name,
                            "default": true
                        }, true)
                    }
                })
                tmpPlayer.on("pause", () => {
                    if (isApp) {
                        recordHistory(1, videoID, episodeID, tmpPlayer.currentTime(), tmpPlayer.remainingTime())
                    }
                })
                tmpPlayer.currentTime(currentTime)
                setPlayer(tmpPlayer)
            } else {
                player.src(url)
                player.off(["pause", "ready"])
                player.currentTime(currentTime)
                player.on("pause", () => {
                    if (isApp) {
                        recordHistory(1, videoID, episodeID, player.currentTime(), player.remainingTime())
                    }
                })
                player.ready(() => {
                    const oldTracks = player.remoteTextTracks();
                    let i = oldTracks.length
                    while (i--) {
                        const item: any = oldTracks[i]
                        player.removeRemoteTextTrack(item);
                    }
                    for (const item of subtitles || []) {
                        player.addRemoteTextTrack({
                            "kind": "subtitles",
                            "src": item.url,
                            "label": item.name,
                            "default": true
                        }, true)
                    }
                })
            }
        }
        return () => {
            if (autoDestroy) {
                player?.dispose()
            }
        }
    }, [videoNode, props, player, url, subtitles,
        autoDestroy, episodeID, isApp, videoID, currentTime]);

    return (
        <div data-vjs-player
            style={{
                width: width, height: height,
                minWidth: minWidth, minHeight: minHeight
            }} >
            <video id={videoKey}
                playsInline
                ref={(node: any) => setVideoNode(node)}
                className={"video-js vjs-big-play-centered " + theme}
                crossOrigin="anonymous"
            />
        </div>
    )
}

export { VideoPlayer }
