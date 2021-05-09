
export interface IBookshelf {
    name: string,
    cover: string,
    layerNum: number,
    partitionNum: number,
}

export interface IUpdateBookshelf {
    id: number,
    name: string,
    cover: string,
    layerNum: number,
    partitionNum: number,
}
