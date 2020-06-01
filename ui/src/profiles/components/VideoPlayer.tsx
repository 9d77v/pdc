import React, { useEffect, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'

import "video.js/dist/video-js.css"
import "./VideoPlayer.less"

export interface SubtitleProps {
    name: string
    url: string
}

export interface VideoPlayerProps {
    episodeID: number
    url: string
    subtitles: [SubtitleProps] | null
    height: number
    width: number
}

const VideoPlayer: React.FC<VideoPlayerProps> = ({
    episodeID,
    url,
    subtitles,
    height,
    width
}) => {

    const videoID = "custom-video" + episodeID
    const [videoNode, setVideoNode] = useState()
    const [player, setPlayer] = useState<VideoJsPlayer>()
    const props: VideoJsPlayerOptions = {
        autoplay: false,
        sources: [{
            src: url,
            type: 'video/mp4',
        }],
        controls: true,
        playbackRates: [0.5, 0.75, 1, 1.25, 1.5, 2],
        loop: true
    };

    useEffect(() => {
        if (videoNode && !player) {
            let tmpPlayer = videojs(videoNode, props, () => {
                for (const item of subtitles || []) {
                    tmpPlayer.addRemoteTextTrack({
                        "kind": "subtitles",
                        "src": item.url,
                        "label": item.name,
                        "default": true
                    }, false)
                }
            })
            setPlayer(tmpPlayer)
        }
        if (videoNode && player) {
            player.src(url)
        }
    }, [videoNode, props, player, url, subtitles]);

    return (
        <div data-vjs-player style={{ width, height }} >
            <video id={videoID}
                ref={(node: any) => setVideoNode(node)}
                className="video-js vjs-big-play-centered"
                crossOrigin="anonymous"
            />
        </div>
    )
}

export { VideoPlayer }