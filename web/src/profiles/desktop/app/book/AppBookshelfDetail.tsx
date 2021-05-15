import { useQuery } from "@apollo/react-hooks"
import { message, Row, Col, Button } from "antd"
import { useEffect, useMemo } from "react"
import { useHistory, useLocation } from "react-router-dom"
import { APP_BOOKSHELF_DETAIL } from "src/gqls/book/bookshelf.query"
import { AppPath } from "src/consts/path"
import { Img } from 'src/components'


export default function AppBookshelfDetail() {
    const location = useLocation()
    const history = useHistory()
    const query = new URLSearchParams(location.search)
    const id = query.get("id")

    const { error, data } = useQuery(APP_BOOKSHELF_DETAIL,
        {
            variables: {
                searchParam: {
                    ids: [id]
                },
                bookPositionSearchParam: {
                },
                bookshelfID: id
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const bookshelf = useMemo(() => {
        return data?.bookshelfs.edges[0]
    }, [data])

    const bookPositionMap = useMemo(() => {
        const m = new Map<string, any>()
        const bps = data?.bookPositions.edges
        if (bps) {
            for (const bp of bps) {
                const key = `${bp.layer}_${bp.partition}`
                let idMap = new Map<number, any>()
                if (m.has(key)) {
                    idMap = m.get(key)
                }
                idMap.set(bp.prevID, bp)
                m.set(key, idMap)
            }
        }
        m.forEach((value: Map<number, any>, key) => {
            let record: any
            value.forEach((v, k) => {
                if (k === 0) {
                    record = v
                }
            })
            let books: any[] = [record]
            while (value.has(record.id)) {
                record = value.get(record.id)
                books.push(record)
            }
            m.set(key, books)
        })
        return m
    }, [data])

    const grid = useMemo(() => {
        if (bookshelf) {
            let rows = []
            for (let i = bookshelf.layerNum; i > 0; i--) {
                let cols = []
                for (let j = 1; j <= bookshelf.partitionNum; j++) {
                    const bps = bookPositionMap.get(`${i}_${j}`) || []
                    let books = []
                    for (const bp of bps) {
                        books.push(<div
                            key={`book_${i}_${j}_${bp.id}`}
                            style={{
                                width: 32, height: "100%", display: "flex", alignItems: "center", justifyContent: "center", background: "black", color: "white", border: "1px solid red", cursor: "pointer",
                                fontSize: 12
                            }}
                            onClick={() => {
                                history.push(AppPath.BOOK_BOOK_DETAIL + "?id=" + bp.bookID)
                            }}
                        >
                            {bp.book?.name}
                        </div>)
                    }
                    cols.push(<Col span={24 / bookshelf.partitionNum} key={`book_col_${i}_${j}`}>
                        <div style={{
                            background: '#663300',
                            height: 130,
                            display: "flex",
                        }}>{books}</div></Col >)
                }
                rows.push(<Row gutter={6} key={`book_row_${i}`} style={{ border: "1px solid white" }}>{cols}</Row>)
            }
            return <div style={{ flex: 1, padding: 6 }}>
                {rows}
            </div>
        }
        return <div style={{ flex: 1 }}></div>
    }, [bookshelf, bookPositionMap, history])
    return (
        <div style={{
            display: 'flex', height: '100%', width: "100%", flexDirection: "column"
        }}>

            <div style={{ padding: 10, fontSize: 24 }}  >
                <Button
                    type="primary"
                    onClick={() => {
                        history.goBack()
                    }}
                    style={{ float: 'left', zIndex: 1, width: 100 }}
                >返回
            </Button>
                <div>{bookshelf?.name}</div> </div>
            <div style={{ display: "flex" }}>
                <div style={{ marginTop: 10, display: 'flex', height: 197 }}>
                    <Img src={bookshelf?.cover} showModal />
                </div>
                {grid}
            </div>
        </div>)
}
