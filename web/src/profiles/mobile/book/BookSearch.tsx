import { useEffect, useMemo, useState } from "react"
import { useHistory } from 'react-router-dom'
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import { Icon, NavBar, SearchBar } from "antd-mobile"
import { MobileBookCard } from "src/profiles/common/book/BookCard"
import { MobileBookshelfCard } from "src/profiles/common/book/BookshelfCard"
import { SEARCH_BOOK } from "src/gqls/book/book.query"


export default function BookSearch() {
    const history = useHistory()
    const [pagination, setPagination] = useState({
        keyword: "",
        page: 1,
        pageSize: 10
    })
    const { error, data } = useQuery(SEARCH_BOOK,
        {
            variables: {
                searchParam: {
                    keyword: pagination.keyword,
                    page: pagination.page,
                    pageSize: pagination.pageSize,
                },
                bookshelfsSearchParam: {
                    page: 1,
                    pageSize: 10,
                }
            },
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const cards = useMemo(() => {
        if (data && data.searchBook.edges) {
            return data.searchBook.edges.map((item: any, index: number) =>
                <MobileBookCard
                    key={index}
                    id={item.id}
                    cover={item.cover}
                    title={item.name}
                />
            )
        }
        return <div></div>
    }, [data])

    const bookshelfs = useMemo(() => {
        if (data && data.bookshelfs.edges) {
            return data.bookshelfs.edges.map((item: any, index: number) =>
                <MobileBookshelfCard
                    key={index}
                    id={item.id}
                    cover={item.cover}
                    title={item.name}
                />
            )
        }
        return <div></div>
    }, [data])

    return (
        <div style={{
            height: "100%",
            overflowY: "scroll"
        }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >书籍索引</NavBar>
            <div style={{ marginTop: 45, display: "flex", flexDirection: "column", width: "100%" }}>
                <SearchBar
                    placeholder="搜索"
                    onSubmit={(value: any) => setPagination({
                        keyword: value,
                        page: 1,
                        pageSize: pagination.pageSize,
                    })}
                    maxLength={8} />
                <div style={{ overflowX: "scroll", display: "flex", height: 152, paddingLeft: 16 }}>
                    {bookshelfs}
                </div>
                <div style={{ flex: 1, width: "100%" }}>
                    {cards}
                </div>
            </div>
        </div>
    )
}
