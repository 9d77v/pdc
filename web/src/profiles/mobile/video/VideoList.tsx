import { useEffect, useMemo, useState } from "react"
import { useHistory } from 'react-router-dom'
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import { Icon, NavBar, SearchBar } from "antd-mobile"
import CheckableTag from "antd/lib/tag/CheckableTag"
import { IVideoPagination } from "src/consts/consts"
import { MobileVideoCard } from "src/profiles/common/video/VideoCard"
import { isMobile } from "src/utils/util"
import { LIST_VIDEO_CARD } from "src/gqls/video/query"


export default function VideoList() {
    const [pagination, setPagination] = useState<IVideoPagination>({
        keyword: "",
        page: 1,
        pageSize: 500,
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
                    isMobile: isMobile()
                }
            },
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const history = useHistory()
    const cards = useMemo(() => {
        if (data && data.searchVideo.edges) {
            const videos = data.searchVideo.edges
            return (videos.map((item: any, index: number) =>
                <MobileVideoCard
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

    const onTagChange = (tag: any, checked: any) => {
        const nextSelectedTags = checked ? [...pagination.selectedTags, tag] : pagination.selectedTags.filter(t => t !== tag)
        setPagination({
            keyword: pagination.keyword,
            page: pagination.page,
            pageSize: pagination.pageSize,
            selectedTags: nextSelectedTags
        })
    }

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
            >视频索引</NavBar>
            <div style={{ marginTop: 45, display: "flex", flexDirection: "column", width: "100%" }}>
                <SearchBar
                    placeholder="搜索"
                    onSubmit={(value: any) => setPagination({
                        keyword: value,
                        page: 1,
                        pageSize: pagination.pageSize,
                        selectedTags: pagination.selectedTags
                    })}
                    maxLength={8} />
                <div >
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
                <div style={{ flex: 1, width: "100%" }}>
                    {cards}
                </div>
            </div>
        </div>
    )
}
