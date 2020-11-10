import React, { useEffect, useState } from "react"
import { useHistory } from 'react-router-dom'
import { message } from "antd"
import "src/styles/button.less"
import { useQuery } from "@apollo/react-hooks"
import { LIST_VIDEO_CARD } from 'src/consts/video.gql'
import Img from "src/components/img"
import { Icon, NavBar, SearchBar } from "antd-mobile"
import CheckableTag from "antd/lib/tag/CheckableTag"
import { IVideoPagination } from "src/consts/consts"
import { AppPath } from "src/consts/path"


export default function VideoList() {
    const [cards, setCards] = useState(<div />)
    const [pagination, setPagination] = useState<IVideoPagination>({
        keyword: "",
        page: 1,
        pageSize: 500,
        selectedTags: []
    })
    const { error, data } = useQuery(LIST_VIDEO_CARD,
        {
            variables: {
                keyword: pagination.keyword,
                page: pagination.page,
                pageSize: pagination.pageSize,
                tags: pagination.selectedTags
            },
            fetchPolicy: "cache-and-network"
        }
    )

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const history = useHistory()
    useEffect(() => {
        if (data && data.searchVideo.edges) {
            const videos = data.searchVideo.edges
            setCards(videos.map((item: any) =>
                <div key={item.id}
                    onClick={() => history.push(AppPath.VIDEO_DETAIL + "?video_id=" + item.id)}
                    style={{
                        width: "30%",
                        margin: "2.5% 0 0 2.5%",
                        height: 210,
                        display: "flex",
                        float: "left",
                        flexDirection: "column"
                    }}
                >
                    <Img src={item.cover} width={"100%"} height={"70%"} />
                    <div style={{
                        fontSize: 12,
                        height: 36,
                        overflow: "hidden",
                        textOverflow: "ellipsis"
                    }}>{item.title}</div>
                    <div style={{ fontSize: 10 }}>全{item.totalNum}话</div>
                </div >
            ))
        }
    }, [data, history])

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
        <div style={{ height: "100%" }}>
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
