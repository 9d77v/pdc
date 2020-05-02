
export interface Subtitle {
    name: string
    url: string
}

export interface Episode {
    id: number
    title: string
    desc: string
    num: number
    cover: string
    url: string
    subtitles: Subtitle[]
}

