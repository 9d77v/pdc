import React, { useEffect, useState } from "react"
import { useHistory } from 'react-router-dom';
import { message, Pagination } from "antd"
import "../../../../style/video.less"
import "../../../../style/button.less"
import { useQuery } from "@apollo/react-hooks";
import { LIST_VIDEO_CARD } from '../../../../consts/video.gql';
import { Img } from "../../../../components/Img";
import Search from "antd/lib/input/Search";
import CheckableTag from "antd/lib/tag/CheckableTag";


interface IPagination {
    keyword: string
    page: number,
    pageSize: number,
    selectedTags: string[]
}

export default function VideoList() {

    const [cards, setCards] = useState(<div />)
    const [pagination, setPagination] = useState<IPagination>({
        keyword: "",
        page: 1,
        pageSize: 12,
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
    );

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
                    onClick={() => history.push('/app/media/videos/' + item.id)}
                    className={"card"}
                >
                    <div style={{ clear: "both" }} />
                    <Img src={item.cover} />
                    <div style={{ marginTop: 5, fontSize: 14 }}>{item.title}</div>
                    <div style={{ fontSize: 12 }}>全{item.totalNum}话</div>
                </div >
            ))
        }
    }, [data, history])

    const onChange = (page: number) => {
        setPagination({
            keyword: pagination.keyword,
            page: page || 1,
            pageSize: pagination.pageSize,
            selectedTags: pagination.selectedTags
        })
    }

    const onTagChange = (tag: any, checked: any) => {
        const nextSelectedTags = checked ? [...pagination.selectedTags, tag] : pagination.selectedTags.filter(t => t !== tag);
        setPagination({
            keyword: pagination.keyword,
            page: pagination.page,
            pageSize: pagination.pageSize,
            selectedTags: nextSelectedTags
        })
    }

    const showTotal = (total: number) => {
        return `共 ${parseInt(((total / pagination.pageSize) + 1).toString())} 页/ ${total} 个`;
    }

    return (
        <div style={{ display: "flex", flexDirection: "column", padding: 12 }}>
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
            <div style={{ marginLeft: 20, marginTop: 6 }}>
                <div className={"pdc-button-selected"}
                    style={{ cursor: "pointer" }}
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
            <div>{cards}</div>
            <Pagination
                showQuickJumper
                onChange={onChange}
                current={pagination.page}
                pageSize={pagination.pageSize}
                total={data?.searchVideo.totalCount}
                showTotal={showTotal}
            />
        </div>
    )
}