import React, { useEffect, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'
import "video.js/dist/video-js.min.css"
import "./index.less"
import "./vjs-theme-lemon.less"

import video_zhCN from 'video.js/dist/lang/zh-CN.json'
const lang: any = video_zhCN
lang["Picture-in-Picture"] = "画中画"
videojs.addLanguage('zh-CN', lang);

export interface LivePlayerProps {
    url: string
    height?: any
    width?: any
    minHeight?: number
    minWidth?: number
    autoplay?: boolean
    autoDestroy?: boolean
}

const LivePlayer: React.FC<LivePlayerProps> = ({
    url,
    height,
    width,
    minWidth,
    minHeight,
    autoplay,
    autoDestroy,
}) => {
    const [videoNode, setVideoNode] = useState()
    const [player, setPlayer] = useState<VideoJsPlayer>()

    const props: VideoJsPlayerOptions = {
        autoplay: autoplay,
        sources: [{
            src: url,
            type: 'application/vnd.apple.mpegurl',
        }],
        language: "zh-CN",
        controls: true,
    };

    if (autoDestroy === undefined) {
        autoDestroy = true
    }
    useEffect(() => {
        if (videoNode && url) {
            if (!player) {
                let tmpPlayer = videojs(videoNode, props)
                setPlayer(tmpPlayer)
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
            <video
                playsInline
                ref={(node: any) => setVideoNode(node)}
                className={"video-js vjs-big-play-centered"}
                crossOrigin="anonymous"
            />
        </div>
    )
}

export { LivePlayer }
