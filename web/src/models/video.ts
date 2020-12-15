
export interface Subtitle {
    name: string
    url: string
}

export interface Episode {
    id: number
    videoID: number
    title: string
    desc: string
    num: number
    cover: string
    url: string
    subtitles: Subtitle[]
}

export interface VideoCardModel {
    id: number
    title: string
    desc: string
    cover: string
    totalNum: number
    episodeID: number
}

export interface IUpdateVideo {
    title: string,
    desc: string,
    cover: string,
    pubDate: number,
    tags: string[],
    isShow: boolean,
    isHideOnMobile: boolean,
    theme: string
}

export const getVideoDetail = (data: any, episodeID: string | null) => {
    const videos = data?.videos.edges
    let video: any
    if (videos?.length > 0) {
        video = videos[0]
    }
    let videoItem = {
        id: 0,
        cover: "",
        title: "",
        desc: "",
        tags: [],
        pubDate: 0,
        episodes: [],
        theme: ""
    }
    if (video) {
        videoItem = {
            id: video.id,
            cover: video.cover,
            title: video.title,
            desc: video.desc,
            episodes: video.episodes,
            theme: video.theme,
            tags: video.tags,
            pubDate: video.pubDate,
        }
    }
    let num = 0
    if (video && video.episodes) {
        for (let i = 0, len = video.episodes.length; i < len; i++) {
            if (Number(video.episodes[i].id) === Number(episodeID)) {
                num = i
                break
            }
        }
    }
    let episodeItem = {
        id: 0,
        url: "",
        subtitles: null
    }
    if (video && video.episodes && num < video.episodes.length) {
        episodeItem = {
            id: video.episodes[num].id,
            url: video.episodes[num].url,
            subtitles: video.episodes[num].subtitles,
        }
    }
    return {
        video: video,
        videoItem: videoItem,
        num: num,
        episodeItem: episodeItem,
        videoSerieses: data?.videoSerieses.edges,
        similarVideos: data?.similarVideos.edges
    }
}