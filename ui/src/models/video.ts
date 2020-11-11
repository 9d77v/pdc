
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
}
