import { useQuery } from "@apollo/react-hooks"
import { message, Descriptions } from "antd"
import { Img } from 'src/components'
import { useEffect, useMemo } from "react"
import { useLocation } from "react-router-dom"
import { BOOK_DETAIL } from "src/gqls/book/book.query"
import dayjs from "dayjs"

export default function BookDetail() {
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const id = query.get("id")

    const { error, data } = useQuery(BOOK_DETAIL,
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
        return data?.books.edges[0]
    }, [data])

    return (
        <div style={{
            display: 'flex', height: '100%', width: "100%", flexDirection: "column", overflowX: "scroll"
        }}>
            <div style={{ marginTop: 10, display: 'flex', flexDirection: "column", height: 197 }}>
                <span style={{ textAlign: "left", paddingLeft: 30, fontSize: 24 }}  >{book?.name}</span>
                <div style={{ height: 160, width: 160 }}><Img src={book?.cover} showModal /></div>
            </div>
            <div style={{ display: "flex" }}>
                <div style={{ width: 500 }}>
                    <Descriptions
                        title="属性"
                        bordered
                        size={"middle"}
                        column={2}
                    >
                        <Descriptions.Item label="书名">{book?.name}</Descriptions.Item>
                        <Descriptions.Item label="ISBN">{book?.isbn}</Descriptions.Item>
                        <Descriptions.Item label="作者">{book?.author}</Descriptions.Item>
                        <Descriptions.Item label="译者">{book?.translator}</Descriptions.Item>
                        <Descriptions.Item label="出版社">{book?.publishingHouse}</Descriptions.Item>
                        <Descriptions.Item label="版次">{book?.edition}</Descriptions.Item>
                        <Descriptions.Item label="印次">{book?.printedTimes}</Descriptions.Item>
                        <Descriptions.Item label="印张">{book?.printedSheets}</Descriptions.Item>
                        <Descriptions.Item label="开本">{book?.format}</Descriptions.Item>
                        <Descriptions.Item label="字数">{book?.wordCount}</Descriptions.Item>
                        <Descriptions.Item label="定价">{book?.pricing}</Descriptions.Item>
                        <Descriptions.Item label="购买价">{book?.purchasePrice}</Descriptions.Item>
                        <Descriptions.Item label="购买时间">{book?.purchaseTime ? dayjs(book.purchaseTime * 1000).format("YYYY-MM-DD HH:mm:ss") : ""}</Descriptions.Item>
                        <Descriptions.Item label="购买途径">{book?.purchaseSource}</Descriptions.Item>
                        <Descriptions.Item label="创建时间">{book?.createdAt ? dayjs(book.createdAt * 1000).format("YYYY-MM-DD HH:mm:ss") : ""}</Descriptions.Item>
                        <Descriptions.Item label="更新时间">{book?.updatedAt ? dayjs(book.updatedAt * 1000).format("YYYY-MM-DD HH:mm:ss") : ""}</Descriptions.Item>

                    </Descriptions>
                </div>

                <div style={{ flex: 1 }}></div>
            </div>
        </div>)
}
