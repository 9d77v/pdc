import { useQuery } from "@apollo/react-hooks"
import { message, Button } from "antd"
import { Img } from 'src/components'
import { useEffect, useMemo } from "react"
import { useHistory, useLocation } from "react-router-dom"
import { APP_BOOK_DETAIL } from "src/gqls/book/book.query"

export default function AppBookDetail() {
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
    const itemStyle = { padding: 2, paddingLeft: 5, width: 300 }
    return (
        <div style={{
            display: 'flex', height: '100%', width: "100%", flexDirection: "column", overflowX: "scroll"
        }}>
            <Button
                type="primary"
                onClick={() => {
                    history.goBack()
                }}
                style={{ float: 'left', marginBottom: 12, marginTop: 5, zIndex: 1, width: 100 }}
            >
                返回
            </Button>
            <div style={{ display: 'flex', height: 260, flexDirection: "column" }}>
                <span style={{ textAlign: "left", paddingLeft: 30, fontSize: 26, fontWeight: 800 }}  >{book?.name}</span>
                <div style={{ marginTop: 10, display: 'flex' }}>
                    <div style={{ height: 160, width: 160 }}><Img src={book?.cover} height={160} width={160} /></div>
                    <div style={{ textAlign: "left", paddingLeft: 50 }}>
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
            </div>

            <div style={{ display: "flex", flexDirection: "column", textAlign: "left", paddingLeft: 20 }}>
                <div style={{ color: "#85dbf5", fontSize: 26, paddingBottom: 10 }}>内容简介:</div>
                <div style={{ width: 600 }}>{book?.desc}
                </div>
            </div>
        </div>)
}
