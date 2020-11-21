import { useQuery } from '@apollo/react-hooks'
import { message } from 'antd'
import { Grid, Icon, NavBar } from 'antd-mobile'
import React, { useEffect, useMemo } from 'react'
import { useHistory } from 'react-router-dom'
import { AppPath } from 'src/consts/path'
import { GET_VIDEO_TAGS } from 'src/consts/video.gql'
import VideoTagSuggest from 'src/profiles/common/video/VideoTagSuggest'
import {
    HistoryOutlined, SearchOutlined
} from '@ant-design/icons'
import { IApp } from 'src/models/app'
import { isMobile } from 'src/utils/util'

const VideoIndex = () => {
    const { error, data } = useQuery(GET_VIDEO_TAGS,
        {
            variables: {
                input: {
                    isMobile: isMobile()
                }
            }
        }
    )
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const cards = useMemo(() => {
        if (data && data.searchVideo.aggResults) {
            const tags = data.searchVideo.aggResults
            const suggests: JSX.Element[] = [
                <VideoTagSuggest
                    key={-1}
                    title={"视频动态"}
                    tag={""}
                    pageSize={6}
                />
            ]
            suggests.push(tags.slice(0, 12).map((tag: any, index: number) => {
                return <VideoTagSuggest
                    key={index}
                    tag={tag.key}
                    pageSize={3}
                />
            }))
            return suggests
        }
    }, [data])

    const pageData: IApp[] = [
        {
            text: "最近播放",
            icon: <HistoryOutlined style={{ fontSize: 26 }} />,
            url: AppPath.HISTORY
        }, {
            text: "视频索引",
            icon: <SearchOutlined style={{ fontSize: 26 }} />,
            url: AppPath.VIDEO_SEARCH
        }
    ]

    const history = useHistory()
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
            >视频推荐</NavBar>
            <div style={{
                display: "flex",
                flexDirection: "column",
                marginTop: 45
            }}>
                <div style={{ padding: 5, marginBottom: 10 }}>
                    <Grid data={pageData}
                        columnNum={2} onClick={(item: any) => history.push(item.url)} />
                </div>
                {cards}
            </div>
        </div>
    )
}

export default VideoIndex