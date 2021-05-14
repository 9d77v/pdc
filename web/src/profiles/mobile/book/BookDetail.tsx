import React, { FC, useEffect, useMemo } from "react"
import { Divider, message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import Img from "src/components/img"
import { useHistory, useLocation } from "react-router-dom"
import { APP_BOOK_DETAIL } from "src/gqls/book/book.query"
import { Icon, NavBar } from "antd-mobile"

const BookDetail: FC = () => {
    const location = useLocation()
    const history = useHistory()
    const query = new URLSearchParams(location.search)
    const id = query.get("id")

    const { error, data } = useQuery(APP_BOOK_DETAIL,
        {
            variables: {
                searchParam: {
                    ids: [id]
                },
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])
    const book = useMemo(() => {
        return data?.searchBook.edges[0]
    }, [data])

    const infoStyle = { padding: 2, paddingLeft: 5, display: "flex" }
    const itemStyle = { padding: 2, paddingLeft: 5, width: 120 }
    return (
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >{book?.name} </NavBar>
            <div style={{ marginTop: 45, display: 'flex', flexDirection: "column", }}>
                <div style={{ padding: 16, display: 'flex' }}>
                    <div style={{ height: 120, width: 120 }}><Img src={book?.cover} height={120} width={120} /></div>
                    <div style={{ textAlign: "left", paddingLeft: 10 }}>
                        <div style={infoStyle}>
                            <span style={itemStyle}>作者：{book?.author}</span>
                            {book?.translator ? <span style={itemStyle}>译者：{book?.translator}</span> : null}
                        </div>
                        <div style={infoStyle}>
                            <span style={itemStyle}>ISBN：{book?.isbn}</span>
                            <span style={itemStyle}>出版社：{book?.publishingHouse}</span>
                        </div>
                        <div style={infoStyle}>
                            <span style={itemStyle}>开本：{book?.format}</span>
                            <span style={itemStyle}>版次：{book?.edition}</span>
                        </div>
                        <div style={infoStyle}>
                            <span style={itemStyle}>印次：{book?.printedTimes}</span>
                            <span style={itemStyle}>印张：{book?.printedSheets}</span>
                        </div>
                        <div style={infoStyle}>
                            <span style={itemStyle}>字数：{book?.wordCount}</span>
                            <span style={itemStyle}>定价：{book?.pricing}</span>
                        </div>
                    </div>
                </div>
                <div style={{ display: "flex", flexDirection: "column", textAlign: "left", paddingLeft: 20 }}>
                    <div style={{ fontSize: 16, paddingBottom: 10, fontWeight: 600 }}>内容简介:</div>
                    <div>{book?.desc}
                    </div>
                </div>
            </div>
        </div>)
}

export default BookDetail
