import { FC, useEffect, useMemo } from "react"
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import Img from "src/components/img"
import { useHistory, useLocation } from "react-router-dom"
import { APP_BOOKSHELF_DETAIL } from "src/gqls/book/bookshelf.query"
import { Icon, NavBar, Tabs } from "antd-mobile"
import { AppPath } from "src/consts/path"

const BookshelfDetail: FC = () => {
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

    const tabs = useMemo(() => {
        let items: any = []
        if (bookshelf?.partitionNum > 0) {
            for (let i = 1; i <= bookshelf?.partitionNum; i++) {
                items.push({ title: '分区' + i })
            }
        }
        return items
    }, [bookshelf])

    const tabContents = useMemo(() => {
        let items: any = []
        if (bookshelf?.partitionNum > 0) {
            for (let i = 1; i <= bookshelf?.partitionNum; i++) {
                let layers: any = []
                for (let j = bookshelf?.layerNum; j >= 1; j--) {
                    const bps = bookPositionMap.get(`${j}_${i}`) || []
                    let books = []
                    for (const bp of bps) {
                        books.push(<div
                            key={`book_${j}_${i}_${bp.id}`}
                            style={{
                                width: 80, padding: 5,
                            }}
                            onClick={() => {
                                history.push(AppPath.BOOK_BOOK_DETAIL + "?id=" + bp.bookID)
                            }}
                        >
                            <div style={{ padding: 5 }}><Img src={bp.book?.cover} height={60} width={60} /></div>
                            <div style={{ fontSize: 8, width: 70, textAlign: "center", textOverflow: "ellipsis" }}>{bp.book?.name}</div>
                        </div>
                        )
                    }
                    layers.push(<div
                        key={`${i}_${j}`}
                        style={{
                            background: '#663300',
                            height: 135,
                            display: "flex",
                            overflowX: "scroll",
                            border: "1px solid white"
                        }}>{books}</div>)
                }
                items.push(
                    <div key={i}
                        style={{ padding: 6, backgroundColor: '#fff', height: 500 }}
                    >
                        {layers}
                    </div>
                )
            }
        }
        return items
    }, [bookshelf, bookPositionMap, history])
    return (
        <div style={{ height: "100%", display: "flex" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >{bookshelf?.name} </NavBar>
            <div style={{ marginTop: 45, display: "flex", width: "100%", flexDirection: "column" }}>
                <Img src={bookshelf?.cover} height={120} width={120} />
                <Tabs tabs={tabs}
                    initialPage={0}
                    animated={true}
                    swipeable={false}
                >
                    {tabContents}
                </Tabs>
            </div>
        </div>)
}

export default BookshelfDetail
