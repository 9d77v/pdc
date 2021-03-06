import { useEffect, useMemo, useState } from "react"
import { message, Pagination } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import Search from "antd/lib/input/Search"
import CheckableTag from "antd/lib/tag/CheckableTag"
import { IVideoPagination } from "src/consts/consts"
import { VideoCard } from "src/profiles/common/video/VideoCard"
import { isMobile } from "src/utils/util"
import { LIST_VIDEO_CARD } from "src/gqls/video/query"


const VideoSearch = () => {
    const [pagination, setPagination] = useState<IVideoPagination>({
        keyword: "",
        page: 1,
        pageSize: 10,
        selectedTags: []
    })
    const { error, data } = useQuery(LIST_VIDEO_CARD,
        {
            variables: {
                searchParam: {
                    keyword: pagination.keyword,
                    page: pagination.page,
                    pageSize: pagination.pageSize,
                    tags: pagination.selectedTags,
                    isMobile: isMobile(),
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
        if (data && data.searchVideo.edges) {
            const videos = data.searchVideo.edges
            return (videos.map((item: any, index: number) =>
                <VideoCard
                    key={index}
                    episodeID={item.episodeID}
                    cover={item.cover}
                    title={item.title}
                    totalNum={item.totalNum}
                />
            ))
        }
        return <div />
    }, [data])

    const onChange = (page: number) => {
        setPagination({
            keyword: pagination.keyword,
            page: page || 1,
            pageSize: pagination.pageSize,
            selectedTags: pagination.selectedTags
        })
    }

    const onTagChange = (tag: any, checked: any) => {
        const nextSelectedTags = checked ? [...pagination.selectedTags, tag] : pagination.selectedTags.filter(t => t !== tag)
        setPagination({
            keyword: pagination.keyword,
            page: pagination.page,
            pageSize: pagination.pageSize,
            selectedTags: nextSelectedTags
        })
    }

    const showTotal = (total: number) => {
        return `共 ${parseInt(((total / pagination.pageSize) + 1).toString())} 页/ ${total} 个`
    }

    return (
        <div style={{ display: "flex", flexDirection: "row", padding: 12 }}>
            <div style={{ display: "flex", flexDirection: "column", width: 1162 }}>
                <div style={{
                    display: "flex",
                    width: 300,
                    margin: "auto",
                    justifyContent: "center",
                    alignItems: "center"
                }}>
                    <Search
                        placeholder="搜索"
                        onSearch={(value: any) => setPagination({
                            keyword: value,
                            page: 1,
                            pageSize: pagination.pageSize,
                            selectedTags: pagination.selectedTags
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
                        total={data?.searchVideo.totalCount}
                        showTotal={showTotal}
                        showSizeChanger={false}
                    />
                </div>
            </div>
            <div style={{ marginLeft: 10, marginTop: 38, display: "flex", flex: 1 }}>
                <div style={{ flex: 1 }}>
                    <div className={"pdc-button-selected"}
                        style={{ cursor: "pointer", width: 66 }}
                        onClick={() => {
                            setPagination({
                                keyword: pagination.keyword,
                                page: pagination.page,
                                pageSize: pagination.pageSize,
                                selectedTags: []
                            })
                        }}>全部</div>
                    <div style={{ maxWidth: 304, minWidth: 233 }}>
                        {data?.searchVideo.aggResults.map((tag: any) => (
                            <CheckableTag
                                className={pagination.selectedTags.indexOf(tag.key) > -1 ? "pdc-button-selected" : "pdc-button"}
                                key={tag.key}
                                checked={pagination.selectedTags.indexOf(tag.key) > -1}
                                onChange={checked => onTagChange(tag.key, checked)}
                            >
                                {tag.key + "(" + tag.value + ")"}
                            </CheckableTag>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    )
}

export default VideoSearch
