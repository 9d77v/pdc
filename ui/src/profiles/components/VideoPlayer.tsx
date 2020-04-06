import React, { useEffect, useState } from 'react'
import videojs, { VideoJsPlayerOptions } from 'video.js'

import "video.js/dist/video-js.css"
export interface SubtitleProps {
    name: String
    url: String
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
    const props: VideoJsPlayerOptions = {
        autoplay: false,
        sources: [{
            src: url,
            type: 'video/mp4'
        }]
    };

    useEffect(() => {
        if (videoNode) {
            videojs(videoNode, props);
        }
    }, [videoNode, props]);

    return (
        <div data-vjs-player style={{ width, height }}>
            <video id={videoID} ref={(node: any) => setVideoNode(node)} controls preload="auto" className="video-js"
                data-setup='{ "playbackRates": [0.5, 1, 1.5, 2,4,8],"loopbutton": true,"language":"zh"  }'
            />
        </div>
    )

}

export { VideoPlayer }