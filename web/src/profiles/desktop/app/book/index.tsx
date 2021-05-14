
import { useQuery } from "@apollo/react-hooks"
import { message, Pagination } from "antd"
import { useEffect, useMemo, useState } from "react"
import { SEARCH_BOOK } from "src/gqls/book/book.query"
import Search from "antd/lib/input/Search"
import { BookCard } from "src/profiles/common/book/BookCard"
import { BookshelfCard } from "src/profiles/common/book/BookshelfCard"

const BookIndex = () => {
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
            const books = data.searchBook.edges
            return books.map((item: any, index: number) =>
                <BookCard
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
            const books = data.bookshelfs.edges
            return books.map((item: any, index: number) =>
                <BookshelfCard
                    key={index}
                    id={item.id}
                    cover={item.cover}
                    title={item.name}
                />
            )
        }
        return <div></div>
    }, [data])

    const onChange = (page: number) => {
        setPagination({
            keyword: pagination.keyword,
            page: page || 1,
            pageSize: pagination.pageSize,
        })
    }


    const showTotal = (total: number) => {
        return `共 ${parseInt(((total / pagination.pageSize) + 1).toString())} 页/ ${total} 个`
    }

    return (
        <div style={{
            display: "flex", padding: 12, background: "#dddddd"
        }}>
            <div style={{ display: "flex", flexDirection: "column", width: 960 }}>
                <div style={{
                    display: "flex",
                    width: 300,
                    height: 100,
                    justifyContent: "center",
                    alignItems: "center",
                    margin: "0 auto",
                }}>
                    <Search
                        placeholder="搜索"
                        onSearch={(value: any) => setPagination({
                            keyword: value,
                            page: 1,
                            pageSize: pagination.pageSize,
                        })}
                        enterButton />
                </div>
                <div>{cards}</div>
                <div>
                    <Pagination
                        style={{ float: "right", marginRight: 20 }}
                        showQuickJumper
                        onChange={onChange}
                        current={pagination.page}
                        pageSize={pagination.pageSize}
                        total={data?.searchBook.totalCount}
                        showTotal={showTotal}
                        showSizeChanger={false}
                    />
                </div>
            </div>
            <div style={{ display: "flex", flex: 1, padding: 20, flexDirection: "column" }}>
                <div>{bookshelfs}</div>
            </div>
        </div>)
}

export default BookIndex
