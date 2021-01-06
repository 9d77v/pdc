import React, { useEffect, useMemo, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'
import "video.js/dist/video-js.min.css"
import "./index.less"
import "./vjs-theme-lemon.less"
import "./vjs-theme-pc.less"


import video_zhCN from 'video.js/dist/lang/zh-CN.json'
import { useLocation } from 'react-router-dom'
import { recordHistory } from 'src/consts/http'
import { isMobile } from 'src/utils/util'

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
    height?: string | number
    width?: string | number
    minHeight?: string | number
    minWidth?: string | number
    maxHeight?: string | number
    maxWidth?: string | number
    autoplay?: boolean
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
    minHeight,
    minWidth,
    maxWidth,
    maxHeight,
    autoplay,
    currentTime = 0,
}) => {
    const location = useLocation();
    const isApp = location.pathname.indexOf("/app") >= 0

    const videoKey = "custom-video" + episodeID
    const [videoNode, setVideoNode] = useState()
    const [player, setPlayer] = useState<VideoJsPlayer>()

    const props: VideoJsPlayerOptions = useMemo(() => {
        return {
            autoplay: autoplay,
            sources: [{
                src: url,
                type: 'video/mp4',
            }],
            language: "zh-CN",
            controls: true,
            playbackRates: [0.5, 1, 2, 4, 16],
            loop: false,
        }
    }, [autoplay, url])

    useEffect(() => {
        let duration = 0
        let timer: any
        const record = () => {
            window.clearInterval(timer);
            if (isApp && player) {
                recordHistory(1, videoID, episodeID, player.currentTime(), player.remainingTime(), duration, Date.now() / 1000)
                duration = 0
            }
        }

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
                tmpPlayer.on('play', () => {
                    timer = setInterval(() => {
                        duration += 0.25
                    }, 250)
                })
                tmpPlayer.on(["pause", "ended"], record)
                tmpPlayer.currentTime(currentTime)
                setPlayer(tmpPlayer)
            } else {
                player.src(url)
                player.off(["pause", "ready", "ended"])
                player.currentTime(currentTime)
                player.on(["pause", "ended"], record)
                player.on('play', () => {
                    timer = setInterval(() => {
                        duration += 0.25
                    }, 250)
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

        window.addEventListener("beforeunload", record, false);
        return () => {
            record()
            window.removeEventListener("beforeunload", record);
            player?.dispose()
        }
    }, [videoNode, props, player, url, subtitles, episodeID, isApp, videoID, currentTime]);

    if (!isMobile()) {
        theme = "vjs-theme-pc " + theme
    }
    return (
        <div data-vjs-player
            style={{
                width: width, height: height,
                minWidth: minWidth, minHeight: minHeight,
                maxWidth: maxWidth, maxHeight: maxHeight
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
