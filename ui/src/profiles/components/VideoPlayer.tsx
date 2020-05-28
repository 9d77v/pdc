import React, { useEffect, useState } from 'react'
import videojs, { VideoJsPlayerOptions, VideoJsPlayer } from 'video.js'

import "video.js/dist/video-js.css"
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
    };

    useEffect(() => {
        if (videoNode) {
            setPlayer(videojs(videoNode, props, () => {
                if (player !== undefined) {
                    if (player.remoteTextTracks().length === 0) {
                        for (const item of subtitles || []) {
                            player.addRemoteTextTrack({
                                "kind": "subtitles",
                                "src": item.url,
                                "label": item.name,
                                "default": true
                            }, false)
                        }
                    }
                }
            }))
        }
    }, [videoNode, props, player, subtitles]);

    useEffect(() => {
        if (videoNode) {
            if (player !== undefined) {
                player.src(url)
            }
        }
    }, [videoNode, player, url]);
    return (
        <div data-vjs-player style={{ width, height }}>
            <video id={videoID} ref={(node: any) => setVideoNode(node)} controls className="video-js"
                data-setup='{ "playbackRates": [0.5, 1, 1.5, 2,4,8,16],"loopbutton": true,"language":"zh"  }'
            />
        </div>
    )
}

export { VideoPlayer }