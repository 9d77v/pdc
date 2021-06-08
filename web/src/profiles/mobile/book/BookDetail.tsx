import { FC, useEffect, useMemo } from "react"
import { message } from "antd"
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
                    <div style={{ textAlign: "left", paddingLeft: 10, display: "flex", flexDirection: "column" }}>
                        <span style={infoStyle}>作者：{book?.author}</span>
                        {book?.translator ? <span style={infoStyle}>译者：{book?.translator}</span> : null}
                        <span style={infoStyle}>ISBN：{book?.isbn}</span>
                        <span style={infoStyle}>出版社：{book?.publishingHouse}</span>
                    </div>
                </div>
                <div style={{ textAlign: "left", paddingLeft: 10 }}>
                    <span style={infoStyle}>开本：{book?.format}</span>
                    <span style={infoStyle}>版次：{book?.edition}</span>
                    <span style={infoStyle}>包装：{book?.packing}</span>
                    <span style={infoStyle}>页数：{book?.pageSize}</span>
                    <span style={infoStyle}>印次：{book?.printedTimes}</span>
                    <span style={infoStyle}>印张：{book?.printedSheets}</span>
                    <span style={infoStyle}>字数：{book?.wordCount} 千字</span>
                    <span style={infoStyle}>定价：{book?.pricing}</span>
                </div>
                <div style={{ display: "flex", flexDirection: "column", textAlign: "left", padding: 10 }}>
                    <div style={{ fontSize: 16, paddingBottom: 10, fontWeight: 600 }}>内容简介:</div>
                    <div>{book?.desc}
                    </div>
                </div>
            </div>
        </div>)
}

export default BookDetail
