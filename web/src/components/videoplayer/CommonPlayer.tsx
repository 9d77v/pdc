import React, { useEffect, useMemo, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'
import "video.js/dist/video-js.min.css"
import "./index.less"

import video_zhCN from 'video.js/dist/lang/zh-CN.json'
import { useLocation } from 'react-router-dom'

const lang: any = video_zhCN
lang["Picture-in-Picture"] = "画中画"
videojs.addLanguage('zh-CN', lang);

export interface VideoPlayerProps {
    id: string | number
    url: string
    height?: any
    width?: any
    minHeight?: number
    minWidth?: number
    autoplay?: boolean
    autoDestroy?: boolean
}

const CommonPlayer: React.FC<VideoPlayerProps> = ({
    id,
    url,
    height,
    width,
    minWidth,
    minHeight,
    autoplay,
    autoDestroy,
}) => {
    const location = useLocation();
    const isApp = location.pathname.indexOf("/app") >= 0

    const videoKey = "custom-video-" + id
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

    if (autoDestroy === undefined) {
        autoDestroy = true
    }

    useEffect(() => {
        if (videoNode && url) {
            if (!player) {
                let tmpPlayer = videojs(videoNode, props, () => { })
                setPlayer(tmpPlayer)
            } else {
                player.src(url)
            }
        }
        return () => {
            if (autoDestroy) {
                player?.dispose()
            }
        }
    }, [videoNode, props, player, url,
        autoDestroy]);

    return (
        <div data-vjs-player
            style={{
                width: width, height: height,
                minWidth: minWidth, minHeight: minHeight
            }} >
            <video id={videoKey}
                playsInline
                ref={(node: any) => setVideoNode(node)}
                className={"video-js vjs-big-play-centered"}
                crossOrigin="anonymous"
            />
        </div>
    )
}

export { CommonPlayer }
