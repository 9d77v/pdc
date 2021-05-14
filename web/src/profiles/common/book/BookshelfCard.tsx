import { useHistory } from "react-router-dom"
import { AppPath } from "src/consts/path"
import React, { FC } from 'react'
import Img from "src/components/img"
import "src/styles/card.less"

interface IBookshelfCardProps {
    id: number
    cover: string
    title: string
}

export const BookshelfCard: FC<IBookshelfCardProps> = ({
    id,
    cover,
    title,
}) => {
    const history = useHistory()
    const link = AppPath.BOOK_BOOKSHELF_DETAIL + "?id=" + id
    return (
        <div
            onClick={() => {
                history.push(link)
            }}
            className={"pdc-bookshelf-card"}
        >
            <div style={{ clear: "both" }} />
            <Img src={cover} height={100} width={120} />
            <div style={{ marginTop: 5, fontSize: 14 }}>{title}</div>
        </div >
    )
}

export const MobileBookshelfCard: React.FC<IBookshelfCardProps> = ({
    id,
    cover,
    title,
}) => {
    const history = useHistory()
    const link = AppPath.BOOK_BOOKSHELF_DETAIL + "?id=" + id
    return (
        <div
            onClick={() => {
                history.push(link)
            }}
            className={"pdc-mobile-bookshelf-card"}
        >
            <div style={{ clear: "both" }} />
            <Img src={cover} height={100} width={120} />
            <div style={{
                marginTop: 5,
                fontSize: 12,
            }}>{title}</div>
        </div >)
}
