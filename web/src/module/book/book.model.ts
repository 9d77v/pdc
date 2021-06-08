import dayjs from 'dayjs'
export interface IBook {
    isbn: string,
    name: string,
    desc: string,
    cover: string,
    author: string[],
    translator: string[],
    publishingHouse: string,
    edition: string,
    printedTimes: string,
    printedSheets: string,
    format: string,
    wordCount: number,
    pricing: number,
    packing: string,
    pageSize: number,
    purchasePrice: number,
    purchaseTime?: dayjs.Dayjs,
    purchaseSource: string,
    bookBorrowUID: number,
}

export interface IUpdateBook {
    id: number,
    isbn: string,
    name: string,
    desc: string,
    cover: string,
    author: string[],
    translator: string[],
    publishingHouse: string,
    edition: string,
    printedTimes: string,
    printedSheets: string,
    format: string,
    wordCount: number,
    pricing: number,
    packing: string,
    pageSize: number,
    purchasePrice: number,
    purchaseTime?: dayjs.Dayjs,
    purchaseSource: string,
    bookBorrowUID: number,
}
