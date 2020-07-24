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
    height?: any
    width?: any
    minHeight?: number
    minWidth?: number
    autoplay?: boolean
    autoDestroy?: boolean
}

const VideoPlayer: React.FC<VideoPlayerProps> = ({
    episodeID,
    url,
    subtitles,
    height,
    width,
    minWidth,
    minHeight,
    autoplay,
    autoDestroy
}) => {

    const videoID = "custom-video" + episodeID
    const [videoNode, setVideoNode] = useState()
    const [player, setPlayer] = useState<VideoJsPlayer>()
    const props: VideoJsPlayerOptions = {
        autoplay: autoplay,
        sources: [{
            src: url,
            type: 'video/mp4',
        }],
        controls: true,
        playbackRates: [0.5, 0.75, 1, 1.25, 1.5, 2],
        loop: false
    };

    if (autoDestroy === undefined) {
        autoDestroy = true
    }
    useEffect(() => {
        if (videoNode && !player && url) {
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
        if (videoNode && player && url) {
            player.src(url)
            var oldTracks = player.remoteTextTracks();
            var i = oldTracks.length;
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
                }, false)
            }
        }
        return () => {
            if (autoDestroy) {
                player?.dispose()
            }
        }
    }, [videoNode, props, player, url, subtitles, autoDestroy]);

    return (
        <div data-vjs-player style={{ width: width, height: height, minWidth: minWidth, minHeight: minHeight }} >
            <video id={videoID}
                ref={(node: any) => setVideoNode(node)}
                className="video-js vjs-big-play-centered"
                crossOrigin="anonymous"
            />
        </div>
    )
}

export { VideoPlayer }
