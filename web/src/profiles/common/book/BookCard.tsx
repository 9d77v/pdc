import { useHistory } from "react-router-dom"
import { AppPath } from "src/consts/path"
import React, { FC } from 'react'
import Img from "src/components/img"
import "src/styles/card.less"

interface IBookCardProps {
    id: number
    cover: string
    title: string
}

export const BookCard: FC<IBookCardProps> = ({
    id,
    cover,
    title,
}) => {
    const history = useHistory()
    const link = AppPath.BOOK_BOOK_DETAIL + "?id=" + id
    return (
        <div
            onClick={() => {
                history.push(link)
            }}
            className={"pdc-book-card"}
        >
            <div style={{ clear: "both" }} />
            <Img src={cover} height={100} width={120} />
            <div style={{ marginTop: 5, fontSize: 14 }}>{title}</div>
        </div >
    )
}

export const MobileBookCard: React.FC<IBookCardProps> = ({
    id,
    cover,
    title,
}) => {
    const history = useHistory()
    const link = AppPath.BOOK_BOOK_DETAIL + "?id=" + id
    return (
        <div
            onClick={() => {
                history.push(link)
            }}
            style={{
                width: "30%",
                margin: "2.5% 0 0 2.5%",
                height: 180,
                display: "flex",
                float: "left",
                flexDirection: "column"
            }}
        >
            <Img src={cover} width={"100%"} height={"70%"} />
            <div style={{
                fontSize: 12,
                height: 36,
                overflow: "hidden",
                textOverflow: "ellipsis"
            }}>{title}</div>
        </div >)
}
