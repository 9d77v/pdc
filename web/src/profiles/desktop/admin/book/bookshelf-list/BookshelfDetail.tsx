import { useMutation, useQuery } from "@apollo/react-hooks"
import { message, Row, Col, Divider, Menu, Dropdown } from "antd"
import { Img } from 'src/components'
import { useEffect, useMemo, useState } from "react"
import { useHistory, useLocation } from "react-router-dom"
import { BOOKSHELF_DETAIL } from "src/gqls/book/bookshelf.query"
import { CREATE_BOOK_POSITION, REMOVE_BOOK_POSITION } from "src/gqls/book/book_position.mutation"
import { AddBookForm } from "./AddBookForm"
import { IBookPosition } from "src/module/book/book_position.model"
import { AdminPath } from "src/consts/path"


export default function BookshelfDetail() {
    const location = useLocation()
    const history = useHistory()
    const query = new URLSearchParams(location.search)
    const id = query.get("id")
    const [addBookData, setAddBookData] = useState<IBookPosition>({
        bookshelf_id: 0,
        layer: 0,
        partition: 0
    })
    const [addBookVisibile, setAddBookVisible] = useState(false)
    const [addBook] = useMutation(CREATE_BOOK_POSITION)
    const [removeBook] = useMutation(REMOVE_BOOK_POSITION)

    const { error, data, refetch } = useQuery(BOOKSHELF_DETAIL,
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

    const onAddBookCreate = async (values: any) => {
        await addBook({
            variables: {
                "input": {
                    "bookshelfID": values.bookshelfID,
                    "bookIDs": values.bookIDs,
                    "layer": values.layer,
                    "partition": values.partition,
                }
            }
        })
        setAddBookVisible(false)
        await refetch()
    }

    const onRemoveBook = async (id: number) => {
        await removeBook({
            variables: {
                "id": id
            }
        })
        await refetch()
    }

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
                        const menu = (
                            <Menu onClick={({ key }: any) => {
                                switch (key) {
                                    case '1': {
                                        history.push(AdminPath.BOOK_DETAIL + "?id=" + bp.bookID)
                                        break
                                    }
                                    case '2': {
                                        onRemoveBook(bp.id)
                                    }
                                }
                            }}>
                                <Menu.Item key="1">详情</Menu.Item>
                                <Menu.Item key="2">移除</Menu.Item>
                            </Menu>)
                        books.push(<Dropdown overlay={menu} key={`book_${i}_${j}_${bp.id}`}><div
                            style={{
                                width: 25, height: "100%", display: "flex", alignItems: "center", justifyContent: "center", background: "black", color: "white", border: "1px solid red", cursor: "pointer",
                                fontSize: 12
                            }}>
                            {bp.book?.name}
                        </div></Dropdown>)
                    }
                    books.push(<div
                        key={`book_${i}_${j}`}
                        style={{ width: 20, height: "100%", display: "flex", alignItems: "center", justifyContent: "center", background: "#3fb4ee", color: "white", cursor: "pointer" }}
                        onClick={
                            () => {
                                setAddBookData({
                                    bookshelf_id: bookshelf?.id,
                                    layer: i,
                                    partition: j
                                })
                                setAddBookVisible(true)
                            }
                        }>
                        放入新书
                    </div>)
                    cols.push(<Col span={24 / bookshelf.partitionNum} key={`book_col_${i}_${j}`}>
                        <div style={{
                            background: '#663300',
                            height: 260,
                            display: "flex",
                        }}>{books}</div></Col >)
                }
                rows.push(<Row gutter={16} key={`book_row_${i}`}>{cols}</Row>,
                    <Divider orientation="center" key={`divider_${i}`}>第{i}层</Divider>)
            }
            return <div style={{ flex: 1 }}>
                {rows}
            </div>
        }
        return <div style={{ flex: 1 }}></div>
    }, [bookshelf, bookPositionMap, history])
    return (
        <div style={{
            display: 'flex', height: '100%', width: "100%", flexDirection: "column", overflowX: "scroll"
        }}>
            <AddBookForm
                visible={addBookVisibile}
                onCreate={onAddBookCreate}
                onCancel={() => {
                    setAddBookVisible(false)
                }}
                addBookData={addBookData}
            />
            <div style={{ marginTop: 10, display: 'flex', flexDirection: "column", height: 197 }}>
                <span style={{ textAlign: "left", paddingLeft: 30, fontSize: 24 }}  >{bookshelf?.name}</span>
                <Img src={bookshelf?.cover} showModal />
            </div>
            {grid}
        </div>)
}
